package impl

import (
	"context"
	"fmt"
)

func (r *repository) Select(ctx context.Context, result interface{}, query string, args ...interface{}) error {
	err := r.sqlConn.SelectContext(ctx, result, query, args...)
	if err != nil {
		return fmt.Errorf("[Postgres][SELECT] %v", err)
	}

	return nil

}
