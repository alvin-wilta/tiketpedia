package impl

import (
	"context"
	"fmt"
)

func (r *repository) Insert(ctx context.Context, query string, args interface{}) error {
	res, err := r.sqlConn.NamedExecContext(ctx, query, args)
	if err != nil {
		return fmt.Errorf("[Postgre][INSERT] %v", err)
	}

	if rowAffected, err2 := res.RowsAffected(); rowAffected <= 0 || err2 != nil {
		return fmt.Errorf("[Postgre][INSERT] No row affected")
	}

	return nil
}
