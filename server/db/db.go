package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	// 打开或创建你SQLite
	db, err := sql.Open("sqlite3", "./zserver.db")
	if err != nil {
		panic(err)
	}

	// 设置全局变量
	DB = db
}

func CreateTables() {
	_, err := DB.Exec(createAuthTableSQL)
	if err != nil {
		panic(err)
	}

}
