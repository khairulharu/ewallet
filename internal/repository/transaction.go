package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type transactionRepository struct {
	db *goqu.Database
}

func NewTransaction(con *sql.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: goqu.New("default", con),
	}
}

func (t transactionRepository) Insert(ctx context.Context, transaction *domain.Transaction) error {
	executor := t.db.Insert("transactions").Rows(goqu.Record{
		"account_id":           transaction.AccountId,
		"sof_number":           transaction.SofNumber,
		"dof_number":           transaction.DofNumber,
		"amount":               transaction.Amount,
		"transaction_type":     transaction.TransactionType,
		"transaction_datetime": transaction.TransactionDatetime,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, transaction)
	return err
}
