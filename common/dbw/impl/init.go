package impl

import (
	"database/sql"

	"github.com/alvin-wilta/tiketpedia/jiraticket/src/common/dbw"
)

type repository struct {
	sqlConn  *sql.DB
	maxRetry int
}

func New(sqlConn *sql.DB, maxRetry int) dbw.Repository {
	return &repository{
		sqlConn:  sqlConn,
		maxRetry: maxRetry,
	}
}
