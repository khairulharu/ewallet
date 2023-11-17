package domain

import (
	"context"

	"github.com/khairulharu/ewallet/dto"
)

type Factor struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	PIN    string `db:"pin"`
}

type FactorRepository interface {
	FindByUser(ctx context.Context, id int64) (Factor, error)
}

type FactorService interface {
	ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error
}
