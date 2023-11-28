package impl

import (
	"context"
	"time"
)

// Redis base set (verb) operations
func (r *repository) Set(ctx context.Context, key string, value interface{}) error {
	_, err := r.redisConn.Set(ctx, key, value, time.Duration(r.ttl)).Result()
	if err != nil {
		return err
	}
	return nil
}
