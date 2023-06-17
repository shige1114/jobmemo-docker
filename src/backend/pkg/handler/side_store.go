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

func (s *SideStore) Side(ctx context.Context, users_id string, sidebar *[]backend.Side) error {
	sql := `
	SELECT c.name AS name, r.companies_id AS id, s.date AS date, r.reject reject, r.offer AS offer ,s.level AS level , s.adjusting AS adjusting
	FROM recruits r 
	JOIN companies c ON r.companies_id = c.id 
	JOIN selections s ON r.companies_id = s.companies_id 
	WHERE r.users_id =  $1
	AND s.level = (SELECT MAX(level) FROM selections WHERE companies_id = r.companies_id);
	`

	rows, err := s.db.QueryxContext(ctx, sql, users_id)
	if err != nil {
		return err
	}

	err = sqlx.StructScan(rows, sidebar)
	if err != nil {
		return err
	}

	return nil

}
