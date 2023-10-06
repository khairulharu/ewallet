package domain

import "context"

type Template struct {
	Code  string `db:"code"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

type TemplateRepository interface {
	FindByCode(ctx context.Context, code string) (Template, error)
}
