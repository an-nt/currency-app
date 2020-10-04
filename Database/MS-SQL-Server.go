package Database

import (
	"database/sql"
	"fmt"

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
		return ms.Db, "Connection failed"
	}
	err = ms.Db.Ping()
	if err != nil {
		return ms.Db, "Connection failed"
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
