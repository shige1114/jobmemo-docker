package mydb

import (
	"database/sql"
	"fmt"
	"reflect"
	"log"

)

type Row struct {
	err error
	rows *sql.Rows
	values *[]interface{}
}
func (r *Row) ScanStruct(dest interface{}) error {
	log.Println("dest:",dest)
	structVal := reflect.ValueOf(dest).Elem()
	structType := structVal.Type()
	defer r.rows.Close()
	cols, err := r.rows.Columns()
	if err != nil {
		return err
	}
	defer r.rows.Close()
	colIdxMap := map[string]int{}
	for i, col := range cols {
		colIdxMap[col] = i
	}
	

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		colName := field.Tag.Get("db")
		if colName == "" {
			colName = field.Name
		}
		_, ok := colIdxMap[colName]
		if !ok {
			continue
		}
		colValPtr := structVal.Field(i).Addr().Interface()
		if err := r.rows.Scan(colValPtr); err != nil {
			return err
		}
	}
	
	return nil
}

func (r *Row) ScanStructs(dest interface{}) error {
	columns, err := r.rows.Columns()
	if err != nil {
		return fmt.Errorf("err:%v", err)
	}

	structType := reflect.TypeOf(dest)
	elementType := structType.Elem()

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = reflect.New(elementType).Interface()
		fmt.Println(values[i])
		err = r.ScanStruct(values[i].(*interface{}))
		if err != nil {
			return err
		}
	}
	r.values = &values

	return nil
}
