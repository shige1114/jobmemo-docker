package mydb

import (
	"context"
	"fmt"
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
func (d *Mysql) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*Row, error) {
	row := d.DB.QueryRowContext(ctx, query, args...)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("")
	}
	return &Row{row: row}, nil
}
func (d *Mysql) ExecContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("エラー %v", err)
	}
	return nil
}

func (d *Mysql) InsertContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("エラー %v", err)
	}
	return nil
}
