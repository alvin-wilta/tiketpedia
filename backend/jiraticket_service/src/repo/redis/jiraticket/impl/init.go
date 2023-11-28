package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/repo/redis/jiraticket"
	"github.com/go-redis/redis"
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
