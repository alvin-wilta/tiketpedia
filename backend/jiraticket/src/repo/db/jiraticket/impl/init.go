package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket/src/repo/db/jiraticket"
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
