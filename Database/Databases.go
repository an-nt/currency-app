package Database

import (
	"CurrencyApp/Model"
	"database/sql"
)

//
type Database struct {
	Db *sql.DB
}

type IDatabaseAccess interface {
	Connect() (*sql.DB, string)
	PostEmployeeRecord(user uint, pass string) error
	GetEmployeeByID(user uint) (Model.Employee, error)
	GetExRateUsdVnd() (uint, error)
	GetPassByID(user uint) (string, error)
	GetFirstRecord(model interface{}) ([]interface{}, error)
}
