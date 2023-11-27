package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/common/dbw"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	service string
	sqlConn *sqlx.DB
}

func New(service string, sqlConn *sqlx.DB) dbw.Repository {
	return &repository{
		service: service,
		sqlConn: sqlConn,
	}
}
