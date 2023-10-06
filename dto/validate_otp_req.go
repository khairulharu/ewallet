package dto

type ValidateOtpReq struct {
	ReferenceID string `json:"reference_id"`
	OTP         string `json:"otp"`
}
