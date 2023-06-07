package handler

import (
	"context"
	"testing"

	"github.com/backend"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestUserHandler(t *testing.T) {
	// insert ユーザ登録
	con, err := sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal("接続エラー", err)
	}
	defer con.DB.Close()
	tx, err := con.Begin()
	if err != nil {
		t.Fatal("トランザクション", err)
	}
	defer tx.Rollback()

	handl := UsersStore{db: con}
	ctx, cancel := context.WithCancel(context.Background())
	uid := uuid.New()
	newUser := backend.Users{
		Id:        uid,
		Name:      "John Doe",
		Email:     "john@example.com",
		Future:    "Some future value",
		Pr:        "Some PR value",
		GoodPoint: "Some good point",
		BadPoint:  "Some bad point",
	}
	err = handl.Insert(ctx, newUser)
	if err != nil {
		t.Fatal(err)
		cancel()
	}
	select_user, err := handl.Users(uid.String(), ctx)
	if err != nil {
		t.Fatal(err)
		cancel()
	}
	t.Log(select_user)

}
