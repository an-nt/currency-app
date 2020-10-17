package Database

import (
	"CurrencyApp/Model"
	"reflect"
)

func (ms *MSSQL) GetExRate() (uint, error) {
	var rate Model.USDVND

	query := "select top 1 * from dbo.USDVND order by Time DESC"
	rows, err := ms.Db.Query(query)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		receiver := []interface{}{&rate.Time, &rate.Exchange}
		err = rows.Scan(receiver...) //pass interface... to pass mutiple variables

		///err = rows.Scan(&rate.Time, &rate.Exchange)
		if err != nil {
			return 0, err
		}
	}
	return rate.Exchange, nil
}

func (ms *MSSQL) GetFirst(model interface{}) ([]interface{}, error) {
	var rate Model.USDVND
	reflecter := reflect.TypeOf(model)
	numFields := reflecter.NumField()
	result := make([]interface{}, numFields)
	result = []interface{}{&rate.Time, &rate.Exchange}
	//var a, b *interface{}
	//result := []interface{}{a, b}

	//query := fmt.Sprintf("SELECT TOP 1 * FROM dbo.%s")
	query := "select top 1 * from dbo.USDVND order by Time DESC"
	rows, err := ms.Db.Query(query)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(result...) //pass interface... to pass mutiple variables

		if err != nil {
			return result, err
		}
	}
	return result, nil
}
