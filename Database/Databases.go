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

func CreateDB(driver, host, user, pass, port, db string) (DB *sql.DB, result string) {
	var adaptor IDatabaseAccess
	switch driver {
	case "sqlserver":
		adaptor = &MSSQL{}
	default:
		result = "No database driver available"
		return DB, result
	}

	DB, result = adaptor.Connect()
	return DB, result
}
