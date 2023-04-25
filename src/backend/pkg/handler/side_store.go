package handler

import (
	"database/sql"
	"fmt"

	"github.com/backend"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx"
)

type SideStore struct {
	db *sql.DB
}

func (s *SideStore) Side(users_id uuid.UUID) ([]backend.Side, error) {
	var side []backend.Side
	sql := `
	SELECT r.companies_id,c.title,
	selections.level,
	selections.adusting,
	selections.date,
	recruits.reject,
	recruits.offer
	FROM recruits r 
	JOIN companies c ON r.companies_id = c.id 
	JOIN selections s ON r.companies_id = s.companies_id 
	WHERE r.users_id = ? ;`
	rows, err := s.db.Query(sql, users_id)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	defer s.db.Close()
	rows.Scan(side)

	return rows, nil

}
