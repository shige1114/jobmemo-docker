package handler

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func NewStore(test bool, uri string) (*Store, error) {
	var db *sqlx.DB
	var err error

	if test != true {
		db, err = sqlx.Open("pgx", uri)
	} else {
		db, err = sqlx.Open("pgx", "postgres://root:password@db:5432/test_db")
	}

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("")
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("")
	}

	return &Store{
		SideStore: &SideStore{db: db},
		MainStore: &MainStore{db: db},
	}, nil
}

type Store struct {
	*SideStore
	*MainStore
	*UsersStore
}
