package redis

import (
	"github.com/alvin-wilta/tiketpedia/backend/core/config"
	"github.com/go-redis/redis"
)

func New(config *config.Config) *redis.Client {
	// opt := Options{
	// 	MaxIdleConn:     10,
	// 	MaxActiveConn:   1000,
	// 	Timeout:         1,
	// 	Wait:            false,
	// 	MaxConnLifetime: 60,
	// }
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       0,
	})
	return client

}
