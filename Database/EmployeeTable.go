package Database

import (
	"CurrencyApp/Model"
	"fmt"
)

type ICheckExist interface {
	CheckExist(user uint, pass string) (Model.Employee, error)
}

func (ms *MSSQL) CheckExist(user uint, pass string) (Model.Employee, error) {
	var emp Model.Employee
	var empIndex []Model.Employee

	query := fmt.Sprintf("SELECT * FROM dbo.Employee WHERE ID = %d AND Password = '%s';", user, pass)
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
		return emp, NoPer()
	case 1:
		return emp, nil
	default:
		return emp, MulPerson()
	}
}

type MulPerErr struct{}

func (e *MulPerErr) Error() string {
	return "Multiple persons with the same ID"
}
func MulPerson() error {
	return &MulPerErr{}
}

type NoPerErr struct{}

func (e *NoPerErr) Error() string {
	return "Unauthorized"
}
func NoPer() error {
	return &NoPerErr{}
}
