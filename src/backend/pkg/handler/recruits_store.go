package handler

import (
	"context"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type RecruitsStore struct {
	db *sqlx.DB
}

func (r *RecruitsStore) Recruits(ctx context.Context, users_id string, companies_id string) (*backend.Recruits, error) {
	var recruit backend.Recruits

	sql := `
	SELECT users_id,companies_id,reject, offer, motivation, good_point, concerns , selections_id
	FROM recruits
	WHERE recruits.users_id = $1 AND recruits.companies_id = $2;
	`
	row := r.db.QueryRowxContext(ctx, sql, users_id, companies_id)

	if err := row.Err(); err != nil {
		return nil, err
	}

	err := row.StructScan(&recruit)
	if err != nil {
		return nil, err
	}

	return &recruit, nil

}

func (r *RecruitsStore) InsertContext(ctx context.Context, recruits backend.Recruits) error {

	sql := `
	INSERT INTO ( users_id,companies_id,reject, offer, motivation, good_point, concerns,selections_id )
	VALUES (:users_id,:companies_id,:reject, :offer, :motivation, :good_point, :concerns, :selections_id)
	`
	_, err := r.db.NamedExecContext(ctx, sql, recruits)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecruitsStore) Update(ctx context.Context, recruit backend.Recruits) error {
	sql := `
	UPDATE recruits SET reject =:reject, offer = :offer ,motivation = :motivation,good_point = :good_point,concerns= :concerns,selections_id = :selections_id
	WHERE users_id = :users_id AND companies_id = :companies_id
	`
	_, err := r.db.NamedExecContext(ctx, sql, recruit)

	if err != nil {
		return err
	}
	return nil
}

func (r *RecruitsStore) Delete(ctx context.Context, users_id, companies_id string) error {
	sql := `
	DELTE FROM recruits WEHERE users_id = $1 AND companies_id =$2;
	`

	_, err := r.db.ExecContext(ctx, sql, users_id, companies_id)
	if err != nil {
		return err
	}
	return nil
}
