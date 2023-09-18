package util

import (
	"errors"

	"github.com/khairulharu/ewallet/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500
	}

}
