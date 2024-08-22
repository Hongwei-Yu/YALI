package relastorage

import "database/sql"

type SqliteConnect struct {
	Servername string
	DBname     string
	Client     *sql.DB
}

type SqliteWriter struct {
	Writer *sql.DB
}
