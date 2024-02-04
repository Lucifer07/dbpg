package ddl

import (
	"database/sql"
)

type DB struct {
	host     string
	port     int
	username string
	password string
	dbName   string
	db       *sql.DB
}
