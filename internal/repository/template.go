package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type templateRepository struct {
	db *goqu.Database
}

func NewTemplate(con *sql.DB) domain.TemplateRepository {
	return &templateRepository{
		db: goqu.New("default", con),
	}
}

func (t templateRepository) FindByCode(ctx context.Context, code string) (template domain.Template, err error) {
	dataset := t.db.From("templates").Where(goqu.Ex{
		"code": code,
	})

	_, err = dataset.ScanStructContext(ctx, &template)
	return
}
