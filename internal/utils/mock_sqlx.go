package utils

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func NewMockSqlx() (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, err
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	return sqlxDb, mock, nil
}
