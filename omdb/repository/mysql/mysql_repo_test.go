package mysql_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	omdbMysqlRepo "github.com/vinbyte/movies/omdb/repository/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestLogRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.Nil(t, err)
		query := "INSERT log SET request=\\? , response=\\?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs("raw request", "raw response").WillReturnResult(sqlmock.NewResult(1, 1))

		or := omdbMysqlRepo.NewOmdbMysqlRepository(db)
		lastID, err := or.LogRequest(context.Background(), "raw request", "raw response")
		assert.Nil(t, err)
		assert.Equal(t, int64(1), lastID)
	})
	t.Run("err prepare context", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.Nil(t, err)
		query := "INSERT log SET request=\\? , response=\\?"
		prep := mock.ExpectPrepare(query).WillReturnError(errors.New("something error when call PrepareContext"))
		assert.NotNil(t, prep)
		or := omdbMysqlRepo.NewOmdbMysqlRepository(db)
		lastID, err := or.LogRequest(context.Background(), "raw request", "raw response")
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), lastID)
	})
	t.Run("err ExecContext", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.Nil(t, err)
		query := "INSERT log SET request=\\? , response=\\?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs("raw request", "raw response").WillReturnError(errors.New("something wrong when call ExecContext"))

		or := omdbMysqlRepo.NewOmdbMysqlRepository(db)
		lastID, err := or.LogRequest(context.Background(), "raw request", "raw response")
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), lastID)
	})
}
