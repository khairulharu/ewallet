package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/ewallet/domain"
)

type loginLogRepository struct {
	db *goqu.Database
}

func NewLoginLog(con *sql.DB) domain.LoginLogRepository {
	return &loginLogRepository{
		db: goqu.New("default", con),
	}
}

func (l loginLogRepository) FindLastAuthorized(ctx context.Context, userId int64) (loginlog domain.LoginLog, err error) {
	dataset := l.db.From("login_log").Where(goqu.Ex{
		"user_id":       userId,
		"is_authorized": true,
	}).Order(goqu.I("id").Desc()).Limit(1)

	if _, err := dataset.ScanStructContext(ctx, &loginlog); err != nil {
		return domain.LoginLog{}, err
	}

	return
}

func (l loginLogRepository) Save(ctx context.Context, login *domain.LoginLog) error {
	executor := l.db.Insert("login_log").Rows(goqu.Record{
		"user_id":       login.UserID,
		"is_authorized": login.IsAuthorized,
		"ip_address":    login.IpAddress,
		"timezone":      login.Timezone,
		"lat":           login.Lat,
		"lon":           login.Lon,
		"access_time":   login.AccessTime,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, login)

	return err
}
