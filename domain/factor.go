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
	Insert(ctx context.Context, factor *Factor) error
	FindByUser(ctx context.Context, id int64) (Factor, error)
}

type FactorService interface {
	CreatePIN(ctx context.Context, req dto.Factor) error
	ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error
}
