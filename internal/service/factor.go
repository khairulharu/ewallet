package service

import (
	"context"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"golang.org/x/crypto/bcrypt"
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

	factor, err := f.factorRepository.FindByUser(ctx, req.UserID)
	if err != nil {
		return err
	}

	if factor == (domain.Factor{}) {
		return domain.ErrValidatePin
	}

	err = bcrypt.CompareHashAndPassword([]byte(factor.PIN), []byte(req.PIN))
	if err != nil {
		return domain.ErrValidatePin
	}

	return nil
}

func (f factorService) CreatePIN(ctx context.Context, req dto.Factor) error {
	pin, _ := bcrypt.GenerateFromPassword([]byte(req.PIN), 12)

	if err := f.factorRepository.Insert(ctx, &domain.Factor{
		PIN:    string(pin),
		UserID: req.UserID,
	}); err != nil {
		return domain.ErrInvalidPin
	}

	return nil
}
