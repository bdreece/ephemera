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
	Begin() (TX, error)
}

type db struct{ *sql.DB }

func (conn *db) Begin() (TX, error) { return conn.DB.Begin() }

func OpenSQLite(dsn DSN) (DB, error) {
	conn, err := sql.Open("sqlite", dsn.String())
	if err != nil {
		return nil, errors.Join(ErrOpenSqlite, err)
	}

	return &db{conn}, nil
}
