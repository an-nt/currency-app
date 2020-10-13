package main

import (
	"CurrencyApp/API"
	"CurrencyApp/Database"
	"CurrencyApp/Server"
	"fmt"
)

var ubuntuhost = "192.168.255.1"
var localhost = "127.0.0.1"

func main() {
	SetUp()
}

func SetUp() {
	ms := Database.MSSQL{}
	db, result := ms.Config(ubuntuhost, "sa", "khtn@2020", "1433", "Supermarket").Connect()
	fmt.Println(result)

	sv := Server.HttpServer{
		Exec: &API.Api{
			DbAccess: &Database.MSSQL{
				Database: Database.Database{Db: db},
			},
		},
	}
	sv.StartServer()
}
