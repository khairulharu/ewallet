package dto

type AuthReq struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
