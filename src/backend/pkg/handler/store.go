package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx"
)

func NewStore(uri string) (*Store, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URI"))

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("")
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("")
	}

	return &Store{
		SideStore: &SideStore{DB: db},
		MainStore: &MainStore{DB: db},
	}, nil
}

type Store struct {
	*SideStore
	*MainStore
}
