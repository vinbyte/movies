package mysql

import (
	"context"
	"database/sql"

	"github.com/vinbyte/movies/domain"

	log "github.com/sirupsen/logrus"
)

type mysqlRepository struct {
	Conn *sql.DB
}

// NewOmdbMysqlRepository will create an object that represent the general.Repository interface
func NewOmdbMysqlRepository(Conn *sql.DB) domain.MysqlOmdbRepository {
	return &mysqlRepository{Conn}
}

func (m *mysqlRepository) LogRequest(ctx context.Context, request string, response string) (int64, error) {
	query := `INSERT log SET request=? , response=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("PrepareContext : ", err)
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, request, response)
	if err != nil {
		log.Error("ExecContext : ", err)
		return 0, err
	}
	lastID, err := res.LastInsertId()
	return lastID, nil
}
