package domain

import (
	"context"

	"github.com/khairulharu/ewallet/dto"
)

type Topup struct {
	ID      string  `db:"id"`
	UserID  int64   `db:"user_id"`
	Amount  float64 `db:"amount"`
	Status  int8    `db:"status"`
	SnapURL string  `db:"snap_url"`
}

type TopupRepository interface {
	FindById(ctx context.Context, id string) (Topup, error)
	Insert(ctx context.Context, t *Topup) error
	Update(ctx context.Context, t *Topup) error
}

type TopupService interface {
	ConfirmedTopup(ctx context.Context, id string) error
	InitializeTopup(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error)
}
