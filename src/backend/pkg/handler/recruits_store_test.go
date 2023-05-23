package handler

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestRecruitStore(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal(err)
	}

	recruits_sotre := RecruitsStore{db: db}

	users_id := "185ffaae-e320-11ed-8886-26359435711c"
	companies_id := "21c39950-e322-11ed-8886-26359435711c"
	ctx := context.Background()
	recruit, err := recruits_sotre.Recruits(ctx, users_id, companies_id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(recruit)
	recruit.Reject = true
	recruit.Offer = true
	fmt.Println(recruit)
	err = recruits_sotre.Update(ctx, *recruit)
	if err != nil {
		t.Fatal(err)
	}
	recruit, err = recruits_sotre.Recruits(ctx, users_id, companies_id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(recruit)

}
