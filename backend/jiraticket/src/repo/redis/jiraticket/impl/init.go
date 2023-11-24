package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket/src/repo/redis/jiraticket"
	"github.com/go-redis/redis/v8"
)

type repository struct {
	redisConn redis.Client
	maxRetry  int
}

func New(redisConn redis.Client, maxRetry int) jiraticket.Repository {
	return &repository{
		redisConn: redisConn,
		maxRetry:  maxRetry,
	}
}
