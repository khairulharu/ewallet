package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/khairulharu/ewallet/dto"
)

type User struct {
	ID                int64        `db:"id"`
	FullName          string       `db:"full_name"`
	Phone             string       `db:"phone"`
	Email             string       `db:"email"`
	Username          string       `db:"username"`
	Password          string       `db:"password"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified_at"`
	EmailVerifiedAt   time.Time    `db:"-"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
	Register(ctx context.Context, req dto.UserRegisterReq) (dto.UserRegisterRes, error)
	ValidateOTP(ctx context.Context, req dto.ValidateOtpReq) error
}
