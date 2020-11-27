package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	httpHandler "github.com/vinbyte/movies/movies/delivery/http"
	omdbRepo "github.com/vinbyte/movies/movies/repository/omdb"
	movieUsecase "github.com/vinbyte/movies/movies/usecase"
	omdbMysqlRepo "github.com/vinbyte/movies/omdb/repository/mysql"
)

func main() {
	_ = godotenv.Load()
	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = "5"
	}
	timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	timeoutContext := time.Duration(timeout) * time.Second
	logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if logMaxSize == 0 {
		logMaxSize = 50 //default 50 megabytes
	}

	lg := &lumberjack.Logger{
		Filename:  "server.log",
		MaxSize:   logMaxSize,
		LocalTime: true,
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(io.MultiWriter(lg, os.Stdout))

	db := initMysql()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.StandardLogger().Writer()
	router := gin.New()
	router.Use(customRequestLogger())
	router.Use(gin.Recovery())

	mor := omdbMysqlRepo.NewOmdbMysqlRepository(db)
	mr := omdbRepo.NewOmdbRepository(timeoutContext)
	mu := movieUsecase.NewMovieUsecase(timeoutContext, mr, mor)
	httpHandler.NewMovieHandler(router, mu)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	router.Run(":" + os.Getenv("PORT"))
}

func initMysql() *sql.DB {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DB")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	connMysql, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(err)
	}
	err = connMysql.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return connMysql
}

func customRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		log.WithFields(log.Fields{
			"ip":       c.ClientIP(),
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"full_url": c.Request.URL.String(),
			"proto":    c.Request.Proto,
			"status":   c.Writer.Status(),
			"latency":  time.Now().Sub(startTime),
			"ua":       c.Request.UserAgent(),
		}).Info()
	}
}
