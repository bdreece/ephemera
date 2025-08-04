//go:generate go tool sqlc generate -f ../../configs/sqlc.yml
package database

import (
	"database/sql"
	"errors"
)

var (
	ErrOpenSqlite = errors.New("database: failed to open sqlite connection")
)

type TX interface {
	DBTX
	Commit() error
	Rollback() error
}

type DB interface {
	DBTX
	Begin() TX
}

type db struct{ *sql.DB }

func (conn *db) Begin() TX { return conn.Begin() }

func OpenSQLite(dsn DSN) (DB, error) {
	conn, err := sql.Open("sqlite", dsn.String())
	if err != nil {
		return nil, errors.Join(ErrOpenSqlite, err)
	}

	return &db{conn}, nil
}
