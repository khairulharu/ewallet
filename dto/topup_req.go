package dto

type TopUpReq struct {
	Amount float64 `json:"amount"`
	UserID int64   `json:"-"`
}
