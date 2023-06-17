package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestMainStore(t *testing.T) {
	users_id := "185ffaae-e320-11ed-8886-26359435711c"
	companies_id := "21c39950-e322-22ed-8886-26359435711c"
	ctx := context.Background()
	var main backend.Main

	db, err := sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal(err)
	}

	store := MainStore{db: db}

	err = store.Main(ctx, users_id, companies_id, &main)
	if err != nil {
		fmt.Println(main)
		t.Fatal(err)
	}

}
