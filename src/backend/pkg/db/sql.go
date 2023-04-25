package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*sql.DB
}

func (d *DB) Query(t string, args ...interface{}) (*Rows, error) {
	rows, err := d.DB.Query(t, args...)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return &Rows{rows}, nil
}

func (d *DB) Exec(t string, args ...interface{}) error {
	_, err := d.DB.Exec(t, args...)
	if err != nil {
		return fmt.Errorf("")
	}
	return nil
}
