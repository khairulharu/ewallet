package domain

import (
	"context"

	"github.com/khairulharu/ewallet/dto"
)

type IpCheckerService interface {
	Query(ctx context.Context, ip string) (dto.IpChecker, error)
}
