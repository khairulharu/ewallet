package domain

import "errors"

var ErrAuthFailed = errors.New("err authentication failed")

var ErrUsernameTaken = errors.New("username allready exist")

var ErrOtpInvalid = errors.New("otp invalid")

var ErrAccountNotFound = errors.New("account not found")

var ErrInquiryNotFound = errors.New("inquiry not found")

var ErrInsuficientBalance = errors.New("insufisiencet balance")

var ErrTransferAccount = errors.New("cannot transfer to your account self")

var ErrValidatePin = errors.New("pin is invalid")
