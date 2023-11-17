package service

import (
	"context"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/internal/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	config config.Midtrans
	envi   midtrans.EnvironmentType
}

func NewMidtrans(cnf *config.Config) domain.MidtransService {
	envi := midtrans.Sandbox
	if cnf.Midtrans.IsProd {
		envi = midtrans.Production
	}

	return &midtransService{
		config: cnf.Midtrans,
		envi:   envi,
	}
}

func (m midtransService) GenerateSnapURL(ctx context.Context, t *domain.Topup) error {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  t.ID,
			GrossAmt: int64(t.Amount),
		},
	}

	var client snap.Client
	client.New(m.config.Key, m.envi)
	snapResp, err := client.CreateTransaction(req)
	if err != nil {
		return err
	}
	t.SnapURL = snapResp.RedirectURL
	return nil
}

func (m midtransService) VerifyPayment(ctx context.Context, orderId string) (bool, error) {
	var client coreapi.Client
	client.New(m.config.Key, m.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	return false, nil
}
