package impl

import (
	"context"
	"fmt"
)

func (r *repository) Delete(ctx context.Context, key string) (int64, error) {
	result, err := r.redisConn.Del(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("[Redis][DEL] %v", err)
	}
	return result, nil
}
