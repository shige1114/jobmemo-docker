package mydb

import (
	"database/sql"
	"fmt"
)

type Mysql struct {
	*sql.DB
	driverName string
}

func NewMysql(db *sql.DB, driverName string) *Mysql {
	return &Mysql{DB: db, driverName: driverName}
}

func Open(driverName, dataSourceName string) (*Mysql, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Mysql{DB: db, driverName: driverName}, nil
}

func (d *Mysql) Query(t string, args ...interface{}) (*Row, error) {
	rows, err := d.DB.Query(t, args...)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return &Row{rows: rows}, nil
}

func (d *Mysql) QueryRow(t string, args ...interface{}) (*Row, error) {
	row := d.DB.QueryRow(t, args...)
	if row.Err() != nil {
		return nil, fmt.Errorf("")
	}
	return &Row{row: row}, nil
}
