package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMid fiber.Handler) {
	h := authApi{
		userService: userService,
	}

	app.Post("token/generate", h.GenerateToken)
	app.Get("token/validate", authMid, h.ValidateToken)
	app.Post("user/register", h.RegisterUser)
	app.Post("user/validate-otp", h.ValidateOTP)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(token)

}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(200).JSON(user)
}

func (a authApi) RegisterUser(ctx *fiber.Ctx) error {
	var req dto.UserRegisterReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	res, err := a.userService.Register(ctx.Context(), req)
	if err != nil {
		// return ctx.SendStatus(util.GetHttpStatus(err))
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(res)
}

func (a authApi) ValidateOTP(ctx *fiber.Ctx) error {
	var req dto.ValidateOtpReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	if err := a.userService.ValidateOTP(ctx.Context(), req); err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON("validate otp succes")
}
