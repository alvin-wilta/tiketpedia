package db

import (
	"log"

	"github.com/alvin-wilta/tiketpedia/backend/core/config"
	"github.com/jmoiron/sqlx"
)

func Init(config *config.Config) *sqlx.DB {
	connString := config.Database.ConnString
	if connString == "" {
		log.Fatalf(connString)
	}
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
