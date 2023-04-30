package mydb

import (
	"context"
)

type QueryContextType interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*Row, error)
}

type ExecContextType interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) error
}

func (d *Mysql) QueryContext(ctx context.Context, query string, args ...interface{}) (*Row, error) {
	row, err := d.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &Row{rows: row}, err
}
func (d *Mysql) ExecContext(ctx context.Context, query string, args ...interface{}) error {

}

func (d *Mysql) InsertContext(ctx context.Context, query string, args interface{}) error {

}
