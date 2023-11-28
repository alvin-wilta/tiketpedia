package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/repo/db/jiraticket"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	dbConn *sqlx.DB
}

func New(
	dbConn *sqlx.DB,
) jiraticket.Repository {
	return &repository{
		dbConn: dbConn,
	}
}
