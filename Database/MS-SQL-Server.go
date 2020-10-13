package Database

import (
	"CurrencyApp/Model"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"

	_ "github.com/denisenkom/go-mssqldb"
)

type MSSQL struct {
	Database
	host   string
	user   string
	pass   string
	port   string
	dbname string
}

func (ms *MSSQL) Connect() (*sql.DB, string) {
	var err error
	conStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", ms.host, ms.user, ms.pass, ms.port, ms.dbname)
	ms.Db, err = sql.Open("sqlserver", conStr)
	if err != nil {
		return ms.Db, err.Error() //"Connection failed"
	}
	err = ms.Db.Ping()
	if err != nil {
		return ms.Db, err.Error() //"Connection failed"
	}
	err = CheckVersion(ms.Db)
	if err != nil {
		return ms.Db, "Version incompatable"
	}

	return ms.Db, "Connection success"
}

func (ms *MSSQL) Config(host, user, pass, port, dbname string) *MSSQL {
	ms = &MSSQL{
		host:   host,
		user:   user,
		pass:   pass,
		port:   port,
		dbname: dbname,
	}
	return ms
}

func CheckVersion(db *sql.DB) error {
	var ver Model.Version

	query := "select top 1 * from dbo.Version order by Time DESC"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&ver.Time, &ver.Version)
		if err != nil {
			return err
		}
	}
	if ver.Version != "1.0" {
		return errors.New("Version incompatable")
		builder, err := ioutil.ReadFile("Version/Version1.0.sql")
		if err != nil {
			return err
		}
		_, err = db.Exec(string(builder))
		if err != nil {
			return err
		}
	}
	return nil
}
