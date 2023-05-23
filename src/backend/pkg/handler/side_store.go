package handler

import (
	"context"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type SideStore struct {
	db *sqlx.DB
}

func (s *SideStore) Side(ctx context.Context, users_id string) (*[]backend.Side, error) {
	side := []backend.Side{}
	sql := `
	SELECT c.id as companies_id, c.name as companies_name, s.level as selections_level, s.date as selections_date, r.reject recruits_reject, r.offer as recruits_offer
FROM recruits r 
JOIN companies c ON r.companies_id = c.id 
JOIN selections s ON r.companies_id = s.companies_id 
WHERE r.users_id = $1  ;`

	rows, err := s.db.QueryxContext(ctx, sql, users_id)
	defer s.db.Close()
	if err != nil {
		return nil, err
	}

	err = sqlx.StructScan(rows, &side)
	if err != nil {
		return nil, err
	}

	return &side, nil

}
