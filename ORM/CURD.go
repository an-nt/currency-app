package ORM

import (
	"CurrencyApp/Database"
)

type ORM struct {
	QueryGenerator QueryGenerate
	DataAccess     Database.Database
}
type IORM interface {
	ReadByPrimaryKey(model interface{}, conditions ...interface{}) (interface{}, error)
}

func (o *ORM) ReadByPrimaryKey(model interface{}, conditions ...interface{}) (interface{}, error) {
	// reflecter := reflect.TypeOf(model)
	// outModel := Model.Model{}

	// PK, err := FindPrimaryKey(model)
	// if err != nil {
	// 	return nil, err
	// }
	// if len(PK) != len(conditions) {
	// 	return nil, errors.New("mismatch primary key and conditions")
	// }
	// query := o.QueryGenerator.GenerateQuery(model, conditions)

	// rows, err := o.DataAccess.Db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	err = rows.Scan(outModel.Fields...)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// fmt.Println(outModel)
	return nil, nil
}

type QueryGenerate interface {
	GenerateQuery(model interface{}, conditions ...interface{}) string
}
