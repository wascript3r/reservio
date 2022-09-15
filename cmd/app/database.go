package main

import (
	"database/sql"

	"github.com/wascript3r/reservio/pkg/repository/pgsql"
)

func openDatabase(driver, connStr string) (*pgsql.Database, error) {
	conn, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return pgsql.NewDatabase(conn), nil
}
