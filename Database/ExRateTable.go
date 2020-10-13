package Database

import "CurrencyApp/Model"

func (ms *MSSQL) GetExRate() (uint, error) {
	var rate Model.USDVND

	query := "select top 1 * from dbo.USDVND order by Time DESC"
	rows, err := ms.Db.Query(query)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&rate.Time, &rate.Exchange)
		if err != nil {
			return 0, err
		}
	}
	return rate.Exchange, nil
}
