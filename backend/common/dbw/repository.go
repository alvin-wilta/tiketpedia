package dbw

import "context"

type Repository interface {
	Select(ctx context.Context, result interface{}, query string, args ...interface{}) error
}
