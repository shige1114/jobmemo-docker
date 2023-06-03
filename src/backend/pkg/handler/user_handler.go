package handler

import (
	"context"
	"fmt"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type UsersStore struct {
	db *sqlx.DB
}

func (s *UsersStore) Users(id string, ctx context.Context) (backend.Users, error) {
	var user backend.Users

	sql := "SELECT id, name, email, future, pr, good_point, bad_point FROM users WHERE id = $1"
	//row := s.db.QueryRowContext(ctx, sql, id)
	err := s.db.GetContext(ctx, &user, sql, id)

	if err != nil {

		return user, err
	}
	return user, nil
}

func (s *UsersStore) Insert(ctx context.Context, user backend.Users) error {
	sql := "INSERT INTO users (id, email, name) VALUES ($1,$2,$3) "

	res, err := s.db.ExecContext(ctx, sql, user.Id, user.Email, user.Name)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersStore) Update(ctx context.Context, user backend.Users) error {
	sql := `UPDATE users
	SET name = :name,
	email = :email, 
	future = :future,
	pr = :pr,
	good_point = :good_point,
	bad_point = :bad_point
	WHERE id = :id
	`
	_, err := s.db.NamedExecContext(ctx, sql, user)

	if err != nil {
		return err
	}
	return nil
}
func (s *UsersStore) Delete(ctx context.Context, id string) error {
	sql := "DELETE FROM users WHERE id = $1 "

	res, err := s.db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
