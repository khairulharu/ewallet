package service

import (
	"context"
	"strconv"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
)

type factorService struct {
	factorRepository domain.FactorRepository
}

func NewFactor(factorRepository domain.FactorRepository) domain.FactorService {
	return &factorService{
		factorRepository: factorRepository,
	}
}

func (f factorService) ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error {
	id, _ := strconv.Atoi(req.UserID)

	factor, err := f.factorRepository.FindByUser(ctx, int64(id))
	if err != nil {
		return err
	}

	if factor.UserID != int64(id) {
		return domain.ErrValidatePin
	}

	return nil
}
