package handler

import (
	"database/sql"
	"fmt"

	"github.com/backend"
	"github.com/google/uuid"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Users(id uuid.UUID) ([]backend.Users, error) {
	var user backend.Users
	sql := `
	SELECT email,name FROM users WHERE id = ?
	`
	row, err := s.db.Query(sql, id)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	defer s.db.Close()

	return row, nil

}
