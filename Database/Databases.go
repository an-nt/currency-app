package Database

import (
	"CurrencyApp/Model"
	"database/sql"
)

type Database struct {
	Db *sql.DB
}

type IDatabaseAccess interface {
	Connect() (*sql.DB, string)
	CheckExist(user uint, pass string) (Model.Employee, error)
	GetExRate() (uint, error)
	GetStoredPassword(user uint) (string, error)
}
