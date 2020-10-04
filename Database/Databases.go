package Database

import (
	"database/sql"
)

type Database struct {
	Db *sql.DB
}

type IDatabaseAccess interface {
	Connect() (*sql.DB, string)
}
