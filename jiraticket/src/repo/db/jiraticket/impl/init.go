package impl

import (
	"github.com/alvin-wilta/tiketpedia/jiraticket/src/common/redisw"
	"github.com/alvin-wilta/tiketpedia/jiraticket/src/repo/redis/jiraticket"
)

type repository struct {
	redis redisw.Repository
}

func New(
	redis redisw.Repository,
) jiraticket.Repository {
	return &repository{
		redis: redis,
	}
}
