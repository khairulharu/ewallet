package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topUpService    domain.TopupService
}

func NewMidtrans(app *fiber.App, midtransService domain.MidtransService, topUpService domain.TopupService) {
	h := &midtransApi{
		midtransService: midtransService,
		topUpService:    topUpService,
	}

	app.Post("midtrans/payment-callback", h.paymentHandlerNotification)
}

func (m midtransApi) paymentHandlerNotification(ctx *fiber.Ctx) error {
	var notificationPayload map[string]interface{}
	if err := ctx.BodyParser(&notificationPayload); err != nil {
		return ctx.SendStatus(400)
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return ctx.SendStatus(400)
	}

	success, _ := m.midtransService.VerifyPayment(ctx.Context(), orderId)
	if success {
		_ = m.topUpService.ConfirmedTopup(ctx.Context(), orderId)
		return ctx.SendStatus(200)
	}
	return ctx.SendStatus(400)
}
