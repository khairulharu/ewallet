package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
)

type transferApi struct {
	transactionService domain.TransactionService
}

func NewTransfer(app *fiber.App, transactionService domain.TransactionService, authMid fiber.Handler) {
	h := transferApi{
		transactionService: transactionService,
	}

	app.Post("transfer/inquiry", authMid, h.TransferInquiry)
	app.Post("transfer/execute", authMid, h.TransferExecute)
}

func (t transferApi) TransferInquiry(ctx *fiber.Ctx) error {
	var req dto.TransferInquiryReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	inquiry, err := t.transactionService.TransferInquiry(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}
	return ctx.Status(200).JSON(inquiry)
}

func (t transferApi) TransferExecute(ctx *fiber.Ctx) error {
	var req dto.TransferExecuteReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	err := t.transactionService.TransferExecute(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetHttpStatus(err)).JSON(dto.Response{
			Message: err.Error(),
		})
	}
	return ctx.SendStatus(200)
}
