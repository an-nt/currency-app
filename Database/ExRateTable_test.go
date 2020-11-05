package Database

import (
	"CurrencyApp/Model"
	"fmt"
	"testing"
)

func TestGetFirstRecord(t *testing.T) {
	ms := MSSQL{}
	db, result := ms.Config("localhost", "sa", "khtn@2020", "1433", "Supermarket").Connect()
	fmt.Println(result)

	var MockModel Model.USDVND
	ms.Db = db
	out, err := ms.GetFirstRecord(MockModel)
	if err != nil {
		t.Errorf("Error due to %s", err.Error())
	}
	fmt.Println(out)
}
