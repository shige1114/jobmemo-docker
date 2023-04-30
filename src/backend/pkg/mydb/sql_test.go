package mydb

import (
	"testing"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestSql(t *testing.T) {

	db, err := Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		panic("databaseエラー")
	}
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		panic("セレクトエラー")
	}
	var user backend.Users
	rows.Scans(&user)
}
