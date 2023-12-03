package storage

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewPostgres(sqlConn string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", sqlConn)
}
