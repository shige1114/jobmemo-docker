package mydb

import (
	"context"
	"testing"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestSql_Context(t *testing.T) {
	db, err := Open("pgx", "postgres://root:password@db:5432/test_db")
	defer db.Close()
	if err != nil {
		t.Fatal("データベースエラー", err)
	}

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, `SELECT c.id, c.name, s.level, s.date, r.reject, r.offer
	FROM recruits r 
	JOIN companies c ON r.companies_id = c.id 
	JOIN selections s ON r.companies_id = s.companies_id 
	WHERE r.users_id = '185ffaae-e320-11ed-8886-26359435711c';`, "21c39950-e322-11ed-8886-26359435711c")
	defer rows.Close()

	if err != nil {
		t.Fatal("取得エラー", err)
	}
	var side backend.Side
	rows.Scans(side)
	t.Log(rows.values)
}
