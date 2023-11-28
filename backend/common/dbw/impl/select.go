package impl

import (
	"context"
)

func (r *repository) Select(ctx context.Context, result interface{}, query string, args ...interface{}) error {
	err := r.sqlConn.SelectContext(ctx, result, query, args...)
	if err != nil {
		return err
	}

	return nil

}
