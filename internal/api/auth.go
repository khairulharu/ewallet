package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type authApi struct {
	userService   domain.UserService
	factorService domain.FactorService
}

func NewAuth(app *fiber.App, userService domain.UserService, factorService domain.FactorService, authMid fiber.Handler) {
	h := authApi{
		userService:   userService,
		factorService: factorService,
	}

	app.Post("token/generate", h.GenerateToken)
	app.Get("token/validate", authMid, h.ValidateToken)
	app.Post("user/register", h.RegisterUser)
	app.Post("user/validate-otp", h.ValidateOTP)
	app.Post("pin/create", authMid, h.RegisterPIN)
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

func (a authApi) RegisterPIN(ctx *fiber.Ctx) error {
	var req dto.PinReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	user := ctx.Locals("x-user").(dto.UserData)

	if err := a.factorService.CreatePIN(ctx.Context(), dto.Factor{
		PIN:    req.PIN,
		UserID: user.ID,
	}); err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}

	return ctx.SendStatus(200)
}
