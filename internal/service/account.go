package service

import (
	"context"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type accountService struct {
	accountRepository domain.AccountRepository
}

func NewAccount(accountRepository domain.AccountRepository) domain.AccountService {
	return &accountService{
		accountRepository: accountRepository,
	}
}

func (a accountService) CreateAccount(ctx context.Context, req dto.AccountReq) error {
	accountNumber := util.GenerateRandomNumber(7)
	if err := a.accountRepository.Insert(ctx, &domain.Account{
		AccountNumber: accountNumber,
		UserId:        req.UserID,
		Balance:       00,
	}); err != nil {
		return domain.ErrCreateAccountInvalid
	}

	return nil
}
