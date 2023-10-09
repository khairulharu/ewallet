package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type repositoryTopup struct {
	db *goqu.Database
}

func NewTopup(con *sql.DB) domain.TopupRepository {
	return &repositoryTopup{
		db: goqu.New("defautl", con),
	}
}

func (r repositoryTopup) FindById(ctx context.Context, id string) (topup domain.Topup, err error) {
	dataset := r.db.From("topup").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &topup)
	return
}

func (r repositoryTopup) Insert(ctx context.Context, t *domain.Topup) error {
	executor := r.db.Insert("topup").Rows(goqu.Record{
		"id":       t.ID,
		"user_id":  t.UserID,
		"amount":   t.Amount,
		"status":   t.Status,
		"snap_url": t.SnapURL,
	}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}

func (r repositoryTopup) Update(ctx context.Context, t *domain.Topup) error {
	executor := r.db.Update("topup").Where(goqu.Ex{
		"id": t.ID,
	}).Set(goqu.Record{
		"amount":   t.Amount,
		"status":   t.Status,
		"snap_url": t.SnapURL,
	}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
