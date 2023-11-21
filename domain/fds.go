package domain

import "context"

type FdsService interface {
	IsAuthorized(ctx context.Context, ip string, userId int64) bool
}
