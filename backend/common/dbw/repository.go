package dbw

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Select()
	// Insert()
	// Update()
	// Delete()
}

func InitPostgres() *sqlx.DB {
	connString := os.Getenv("POSTGRES_CONN_STRING")
	if connString == "" {
		log.Fatalf(connString)
	}
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
