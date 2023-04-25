package db

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Rows struct {
	*sql.Rows
}

func (r *Rows) ScanStruct(dest *interface{}) error {
	structVal := reflect.ValueOf(dest).Elem()
	structType := structVal.Type()

	cols, err := r.Columns()
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
		if err := r.Scan(colValPtr); err != nil {
			return err
		}
	}

	return r.Err()
}

func (r *Rows) ScanStructs(dest *[]interface{}) error {
	for r.Next() {
		if err := r.ScanStruct(dest); err != nil {
			return fmt.Errorf("")
		}
	}
	defer r.Close()
	return r.Err()
}
