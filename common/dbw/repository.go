package dbw

import (
	"database/sql"
	"log"
	"os"
)

type Repository interface {
	Select()
	Insert()
	Update()
	Delete()
}

func InitPostgres() *sql.DB {
	connString := os.Getenv("POSTGRES_CONN_STRING")
	if connString == "" {
		log.Fatalf(connString)
	}
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
