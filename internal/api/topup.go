package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type topUpApi struct {
	topUpService domain.TopupService
}

func NewTopUp(app *fiber.App, authMid fiber.Handler, topUpService domain.TopupService) {
	h := &topUpApi{
		topUpService: topUpService,
	}

	app.Post("topup/initalize", authMid, h.InitializeTopup)
}

func (t topUpApi) InitializeTopup(ctx *fiber.Ctx) error {
	var req dto.TopUpReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	user := ctx.Locals("x-user").(dto.UserData)
	req.UserID = user.ID

	res, err := t.topUpService.InitializeTopup(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(res)
}
