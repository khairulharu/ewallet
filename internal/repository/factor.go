package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type factorRepository struct {
	db *goqu.Database
}

func NewFactor(con *sql.DB) domain.FactorRepository {
	return &factorRepository{
		db: goqu.New("default", con),
	}
}

func (f factorRepository) FindByUser(ctx context.Context, id int64) (factor domain.Factor, err error) {
	dataset := f.db.From("factors").Where(goqu.Ex{
		"user_id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &factor)
	return
}

func (f factorRepository) Insert(ctx context.Context, factor *domain.Factor) error {
	executor := f.db.Insert("factors").Rows(goqu.Ex{
		"user_id": factor.UserID,
		"pin":     factor.PIN,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, factor)

	return err
}
