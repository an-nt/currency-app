package Database

import (
	"CurrencyApp/Model"
	"errors"
	"fmt"
)

func (ms *MSSQL) PostEmployeeRecord(user uint, pass string) error {
	query := fmt.Sprintf(`INSERT dbo.Employee (ID, Password)
						VALUES ("%s", "%s");`, user, pass)
	_, err := ms.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MSSQL) GetEmployeeByID(user uint) (Model.Employee, error) {
	var emp Model.Employee
	var empIndex []Model.Employee

	query := fmt.Sprintf("SELECT * FROM dbo.Employee WHERE ID = %d;", user)
	rows, err := ms.Db.Query(query)
	if err != nil {
		return emp, err
	}

	for rows.Next() {
		err = rows.Scan(&emp.ID, &emp.FullName, &emp.Male, &emp.Nationality, &emp.Password, &emp.DirectManager)
		if err != nil {
			fmt.Println(err.Error()) //Maybe have a minor error
		}
		empIndex = append(empIndex, emp)
	}
	switch len(empIndex) {
	case 0:
		return emp, errors.New("Unauthorized")
	case 1:
		return emp, nil
	default:
		return emp, errors.New("Multiple results")
	}
}

func (ms *MSSQL) GetPassByID(user uint) (string, error) {
	var HashedPass string
	var PassArray []string
	query := fmt.Sprintf("SELECT Password FROM dbo.Employee WHERE ID = %d", user)
	rows, err := ms.Db.Query(query)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err = rows.Scan(&HashedPass)
		if err != nil {
			return "", err
		}
		PassArray = append(PassArray, HashedPass)
	}
	switch len(PassArray) {
	case 0:
		return "", errors.New("Unauthorized")
	case 1:
		return HashedPass, nil
	default:
		return "", errors.New("Multiple results")
	}
}
