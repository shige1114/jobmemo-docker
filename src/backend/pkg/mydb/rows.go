package mydb

import (
	"database/sql"
	"log"
	"reflect"
)

type Row struct {
	err    error
	rows   *sql.Rows
	row    *sql.Row
	values interface{}
	value  interface{}
}

func (r *Row) ScanStruct(dest interface{}) error {
	structVal := reflect.ValueOf(dest).Elem()
	structType := structVal.Type()

	defer r.rows.Close()
	cols, err := r.rows.Columns()
	if err != nil {
		return err
	}

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

func (r *Row) Scan(dest interface{}) error {

	structVal := reflect.ValueOf(dest).Elem()
	structType := structVal.Type()

	r.value = reflect.New(structType).Interface()
	err := r.ScanStruct(r.value)

	return err
}

func (r *Row) Scans(dest interface{}) error {
	values := []interface{}{}
	structVal := reflect.ValueOf(dest).Elem()
	structType := structVal.Type()
	log.Println(structType)
	for i := 0; r.rows.Next(); i++ {
		value := reflect.New(structType).Interface()
		values = append(values, value)
		err := r.ScanStruct(values[i])
		if err != nil {
			return err
		}
	}
	r.values = &values

	return nil
}
