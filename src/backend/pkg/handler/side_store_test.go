package handler

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestSideStore(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal(err)
	}

	side_store := SideStore{db: db}
	ctx := context.Background()
	side, err := side_store.Side(ctx, "185ffaae-e320-11ed-8886-26359435711c")
	if err != nil {
		t.Fatal(err)
	}
	defer fmt.Print(side)
}
