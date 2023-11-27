package impl

import (
	"context"
	"fmt"
)

func (r *repository) Exists(ctx context.Context, key string) (int64, error) {
	result, err := r.redisConn.Exists(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("[Redis][EXIST] %v", err)
	}
	return result, nil
}
