package mydb

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*sql.DB
	driverName string
}

func NewDB (db *sql.DB,driverName string) *DB {
	return &DB{DB: db, driverName: driverName}
}

func (d *DB) Open(driverName, dataSourceName string)(*DB ,error){
	db, err:=sql.Open(driverName,dataSourceName)
	if err != nil {
		return nil ,err
	}
	return &DB{DB:db,driverName:driverName},nil
}

func (d *DB) Query(t string, args ...interface{}) (*Row, error) {
	rows, err := d.DB.Query(t, args...)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return &Row{rows:rows}, nil
}


