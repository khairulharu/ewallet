package dto

type TransferInquiryReq struct {
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}
