package impl

import (
	"github.com/alvin-wilta/tiketpedia/jiraticket/src/common/redisw"
	"github.com/go-redis/redis/v8"
)

type repository struct {
	redisConn redis.Client
	maxRetry  int
}

func New(redisConn redis.Client, maxRetry int) redisw.Repository {
	return &repository{
		redisConn: redisConn,
		maxRetry:  maxRetry,
	}
}
