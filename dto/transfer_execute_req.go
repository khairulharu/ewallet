package dto

type TransferExecuteReq struct {
	InquiryKey string `json:"inquiry_key"`
	PIN        string `json:"pin"`
}
