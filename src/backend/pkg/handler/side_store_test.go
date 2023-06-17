package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestSidebarStore(t *testing.T) {
	users_id := "185ffaae-e320-11ed-8886-26359435711c"
	ctx := context.Background()
	var sidebar []backend.Side

	db, err := sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal(err)
	}

	store := SideStore{db: db}

	err = store.Side(ctx, users_id, &sidebar)
	if err != nil {
		fmt.Println(sidebar)
		t.Fatal(err)
	}

}
