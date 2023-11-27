package redisw

import "context"

type Repository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, key string) (int64, error)
}
