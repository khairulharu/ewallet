package util

import (
	"errors"

	"github.com/khairulharu/ewallet/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.ErrUsernameTaken):
		return 400
	case errors.Is(err, domain.ErrOtpInvalid):
		return 400
	case errors.Is(err, domain.ErrValidatePin):
		return 400
	case errors.Is(err, domain.ErrInvalidPin):
		return 400
	default:
		return 500
	}

}
