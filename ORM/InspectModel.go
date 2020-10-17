package ORM

import (
	"errors"
	"reflect"
)

func FindPrimaryKey(model interface{}) ([]reflect.StructField, error) {
	var PrimaryKey []reflect.StructField

	reflecter := reflect.TypeOf(model)
	if reflecter.NumField() == 0 {
		return PrimaryKey, errors.New("Model with no field")
	}

	for i := 0; i < reflecter.NumField(); i++ {
		value, ok := reflecter.Field(i).Tag.Lookup("MyORM")
		if ok && value == "ID" {
			PrimaryKey = append(PrimaryKey, reflecter.Field(i))
		}
	}

	if len(PrimaryKey) == 0 {
		return PrimaryKey, errors.New("No primary key")
	}
	return PrimaryKey, nil
}

func ListField(model interface{}) ([]interface{}, error) {
	return nil, nil
}
