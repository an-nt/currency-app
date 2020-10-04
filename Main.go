package main

import (
	"CurrencyApp/Database"
	"fmt"
)

func main() {
	ms := Database.MSSQL{}
	_, result := ms.Config("localhost", "sa", "khtn@2020", "1433", "Supermarket").Connect()
	fmt.Println(result)
}
