package mydb

import (
	"database/sql"
	"log"
	"testing"

	"github.com/backend"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestScanStruct(t *testing.T) {

	con, err := sql.Open("pgx", "postgres://root:password@db:5432/test_db")
	var user backend.Users
	if err != nil {
		t.Fatalf("接続エラー:%v", err)
	}
	defer con.Close()

	log.Println("user;", user)

	row, err := con.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("データベースエラー：%v", err)
	}

	rows := Row{rows: row}

	log.Println("row:", rows)

	if err := rows.Scans(&user); err != nil {
		log.Printf("%v", err)
	}

}
