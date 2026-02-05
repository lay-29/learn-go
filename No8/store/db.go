package store

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://postgres:lianzeyu29@localhost:54326/demo_test")
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("db ping failed:", err)
	}

	log.Println("PostgreSQL connected")
}
