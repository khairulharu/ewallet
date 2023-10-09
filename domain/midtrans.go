package domain

import "context"

type MidtransService interface {
	GenerateSnapURL(ctx context.Context, t *Topup) error
	VerifyPayment(ctx context.Context, data map[string]interface{}) (bool, error)
}
