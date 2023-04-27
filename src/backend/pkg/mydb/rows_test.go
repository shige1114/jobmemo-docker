package mydb

import (
	"testing"
	"database/sql"
	"github.com/backend"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"

)

func TestScanStruct(t *testing.T) {

	con, err := sql.Open("pgx", "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatalf("接続エラー:%v",err)
	}
	defer con.Close()

	var user backend.Users

	log.Println(user)

	row,err := con.Query("SELECT * FROM users LIMIT 1")
	if err != nil {
		fmt.Errorf("データベースエラー：%v",err)
	}
	rows := Row{rows:row}

	
	if err := rows.ScanStruct(user); err != nil {
		fmt.Errorf("スキャンエラー：%v",err)
	}

}
