package domain

import (
	"context"

	"github.com/khairulharu/ewallet/dto"
)

type Account struct {
	ID            int64   `db:"id"`
	UserId        int64   `db:"user_id"`
	AccountNumber string  `db:"account_number"`
	Balance       float64 `db:"balance"`
}

type AccountRepository interface {
	Insert(ctx context.Context, account *Account) error
	FindByUserID(ctx context.Context, id int64) (Account, error)
	FindByAccountNumber(ctx context.Context, accNumber string) (Account, error)
	Update(ctx context.Context, account *Account) error
}

type AccountService interface {
	CreateAccount(ctx context.Context, req dto.AccountReq) error
}
