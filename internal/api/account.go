package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type accountApi struct {
	accountService domain.AccountService
}

func NewAccount(app *fiber.App, authMid fiber.Handler, accountService domain.AccountService) {
	h := accountApi{
		accountService: accountService,
	}

	app.Post("account/create", authMid, h.CreateAccount)
}

func (a accountApi) CreateAccount(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user").(dto.UserData)
	var accountID = dto.AccountReq{
		UserID: user.ID,
	}
	if err := a.accountService.CreateAccount(ctx.Context(), accountID); err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}

	return ctx.SendStatus(200)
}
