package handler

import (
	"context"
	"fmt"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type MainStore struct {
	db *sqlx.DB
}

type MainStoreInterface interface {
	Main(ctx context.Context, users_id, companies_id string, main *backend.Main) error
}
type Tmp struct {
	Level int ``
	Types int ``
}

func (s *MainStore) Main(ctx context.Context, users_id, companies_id string, main *backend.Main) error {
	var sql1 string
	var sql2 string
	var tmp = make(map[int]int)
	sql1 = `
	SELECT c.name AS name,s.pass AS pass,s.fail AS s.fail, s.date AS date, r.reject reject, r.offer AS offer ,s.level AS level ,s.type AS type, s.adjusting AS adjusting, r.motivation AS motivation,r.good_point AS good_point,r.concerns AS concerns
	FROM recruits r 
	JOIN companies c ON r.companies_id = c.id 
	JOIN selections s ON r.companies_id = s.companies_id 
	WHERE r.users_id =  $1 AND r.companies_id = $2
	AND s.level = (SELECT MAX(level) FROM selections WHERE companies_id = r.companies_id);
	`

	row := s.db.QueryRowxContext(ctx, sql1, users_id, companies_id)
	if row.Err() != nil {
		fmt.Println(1)
		return row.Err()
	}

	err := row.StructScan(main)
	if err != nil {
		fmt.Println(2)
		return err
	}
	sql2 = `
	SELECT s.level AS level,s.type AS type
	FROM selections s
	WHERE s.users_id =  $1 AND s.companies_id = $2;
	`
	rows, err := s.db.QueryxContext(ctx, sql2, users_id, companies_id)
	if err != nil {
		fmt.Println(3)
		return err
	}
	err = ScanMap(rows, &tmp)

	main.Selections = tmp

	return nil
}

func ScanMap(rows *sqlx.Rows, tmp *map[int]int) error {
	var level, types int
	for rows.Next() {
		err := rows.Scan(&level, &types)
		if err != nil {
			fmt.Println(4)
			return err
		}
		(*tmp)[level] = types
	}

	if err := rows.Err(); err != nil {
		fmt.Println(5)
		return err
	}

	return nil
}
