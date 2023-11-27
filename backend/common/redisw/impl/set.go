package impl

import (
	"context"
	"fmt"
	"time"
)

// Redis base set (verb) operations
func (r *repository) Set(ctx context.Context, key string, value interface{}) error {
	err := r.redisConn.Set(ctx, key, value, time.Duration(r.ttl))
	if err != nil {
		return fmt.Errorf("[Redis][SET] Error %v", err)
	}
	return nil
}
