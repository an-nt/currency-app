package API

import (
	"CurrencyApp/Model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	signUpFailed  = "Sign-up failed"
	signUpSuccess = "Sign-up successed"
)

func (a *Api) SignUp(user uint, pass string) (string, error) {
	var err error
	var emp Model.Employee
	var hash []byte

	emp, err = a.DbAccess.GetEmployeeByID(user)
	if err != nil {
		return signUpFailed, err
	}
	if emp.ID == user {
		return signUpFailed, errors.New("Member ID has been existed")
	}
	hash, err = bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return signUpFailed, err
	}
	err = a.DbAccess.PostEmployeeRecord(user, string(hash))
	if err != nil {
		return signUpFailed, err
	}
	return signUpSuccess, nil
}
