package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL用のドライバ
)

func main() {
	// データベースへの接続情報
	connStr := "user=myuser dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベースへの接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// SQLクエリの実行
	rows, err := db.Query("SELECT id, name FROM mytable")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 結果の処理
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// エラーチェック
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
