package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/todo_list")
	if err != nil {
		log.Fatalf("error opning with database %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("error ping with database %v", err)
	}
}
