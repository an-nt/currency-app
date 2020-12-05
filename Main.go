package main

import (
	"CurrencyApp/API"
	"CurrencyApp/Database"
	"CurrencyApp/Server"
	"fmt"
)

var ubuntuhost = "192.168.255.1"
var localhost = "127.0.0.1"
var containerhost = "172.18.0.2"

func main() {
	SetUp()
}

func SetUp() {
	ms := Database.MSSQL{}
	db, result := ms.Config(containerhost, "sa", "khtn@2020", "1234", "Supermarket").Connect()

	fmt.Println(result)

	mssql := &Database.MSSQL{}
	mssql.Db = db
	sv := Server.HttpServer{
		Exec: &API.Api{
			DbAccess: mssql,
		},
	}
	sv.StartServer()
}
