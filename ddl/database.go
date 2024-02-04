package ddl

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	host     string
	port     int
	username string
	password string
	dbName   string
	db       *sql.DB
}
