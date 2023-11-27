package impl

import (
	"github.com/alvin-wilta/tiketpedia/backend/common/redisw"
	"github.com/go-redis/redis/v8"
)

type repository struct {
	redisConn *redis.Client
	maxRetry  int
	ttl       int
}

// Redis base initialization, atomic operations to be used on all services
func New(redisConn *redis.Client, maxRetry int) redisw.Repository {
	return &repository{
		redisConn: redisConn,
		maxRetry:  maxRetry,
	}
}
