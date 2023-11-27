package impl

import (
	"context"
)

// Redis base get operations
func (r *repository) Get(ctx context.Context, key string) (string, error) {
	result, err := r.redisConn.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
