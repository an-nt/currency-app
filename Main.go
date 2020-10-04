package main

import (
	"CurrencyApp/API"
	"CurrencyApp/Database"
	"CurrencyApp/Server"
	"fmt"
)

func main() {
	ms := Database.MSSQL{}
	db, result := ms.Config("localhost", "sa", "khtn@2020", "1433", "Supermarket").Connect()
	fmt.Println(result)

	sv := Server.HttpServer{
		LoginFormartter: &Server.LoginFormatter{
			Handler: &API.ExecLogin{
				Checker: &Database.MSSQL{
					Database: Database.Database{db},
				},
			},
		},
		GetExRateFormatter: &Server.GetExRateFormatter{
			Handler: &API.ExecGetExRate{
				Finder: &Database.MSSQL{
					Database: Database.Database{db},
				},
			},
		},
	}
	sv.StartServer()
}
