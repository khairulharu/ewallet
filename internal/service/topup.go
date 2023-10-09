package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
)

type topupService struct {
	notificationService domain.NotificationService
	midtransService     domain.MidtransService
	topUpRepository     domain.TopupRepository
	accountRepository   domain.AccountRepository
}

func NewTopup(notificationService domain.NotificationService, midatransService domain.MidtransService, topUpRepository domain.TopupRepository, accountRepository domain.AccountRepository) domain.TopupService {
	return &topupService{
		notificationService: notificationService,
		midtransService:     midatransService,
		topUpRepository:     topUpRepository,
		accountRepository:   accountRepository,
	}
}

func (t topupService) InitializeTopup(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error) {
	topUp := domain.Topup{
		ID:     uuid.NewString(),
		UserID: req.UserID,
		Status: 0,
		Amount: req.Amount,
	}

	err := t.midtransService.GenerateSnapURL(ctx, &topUp)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	err = t.topUpRepository.Insert(ctx, &topUp)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	return dto.TopUpRes{
		SnapURL: topUp.SnapURL,
	}, nil
}

func (t topupService) ConfirmedTopup(ctx context.Context, id string) error {
	topUp, err := t.topUpRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if topUp == (domain.Topup{}) {
		return errors.New("errors topUp not found")
	}

	account, err := t.accountRepository.FindByUserID(ctx, topUp.UserID)
	if err != nil {
		return err
	}

	if account == (domain.Account{}) {
		return domain.ErrAccountNotFound
	}

	account.Balance += topUp.Amount
	err = t.accountRepository.Update(ctx, &account)
	if err != nil {
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%2.f", topUp.Amount),
	}

	err = t.notificationService.Insert(ctx, "TOPUP_SUCCES", account.UserId, data)
	if err != nil {
		return err
	}

	return nil
}
