package impl

import (
	"context"
)

func (r *repository) Exists(ctx context.Context, key string) (int64, error) {
	result, err := r.redisConn.Exists(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
