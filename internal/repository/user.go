package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		goqu.New("postgres", con),
	}
}

func (u userRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) Insert(ctx context.Context, user *domain.User) error {
	executor := u.db.Insert("users").Rows(goqu.Record{
		"full_name": user.FullName,
		"phone":     user.Phone,
		"username":  user.Username,
		"password":  user.Password,
		"email":     user.Email,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, user)
	return err
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	user.EmailVerifiedAtDB = sql.NullTime{
		Time:  user.EmailVerifiedAt,
		Valid: true,
	}

	executor := u.db.Update("users").Where(goqu.Ex{
		"id": user.ID,
	}).Set(goqu.Record{
		"full_name":         user.FullName,
		"phone":             user.Phone,
		"username":          user.Username,
		"password":          user.Password,
		"email":             user.Email,
		"email_verified_at": user.EmailVerifiedAtDB,
	}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
