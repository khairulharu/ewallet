package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/api"
	"github.com/khairulharu/ewallet/internal/component"
	"github.com/khairulharu/ewallet/internal/config"
	"github.com/khairulharu/ewallet/internal/middleware"
	"github.com/khairulharu/ewallet/internal/repository"
	"github.com/khairulharu/ewallet/internal/service"
	"github.com/khairulharu/ewallet/internal/sse"
)

func main() {
	config := config.Get()
	dbConnection := component.GetDatabaseConnection(config)
	cacheConnection := repository.NewRedisClient(config)

	hub := &dto.Hub{
		NotificationChannel: map[int64]chan dto.NotificationData{},
	}

	userRepository := repository.NewUser(dbConnection)
	accontRepository := repository.NewAccount(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)
	notificationRepository := repository.NewNotification(dbConnection)
	templateRepository := repository.NewTemplate(dbConnection)
	topUpRepository := repository.NewTopup(dbConnection)

	notificationService := service.NewNotification(notificationRepository, templateRepository, hub)
	emailService := service.NewEmail(config)
	userService := service.NewUser(userRepository, cacheConnection, emailService)
	transactionService := service.NewTransaction(accontRepository, transactionRepository, cacheConnection, notificationService)
	midtransService := service.NewMidtrans(config)
	topupService := service.NewTopup(notificationService, midtransService, topUpRepository, accontRepository, transactionRepository)

	authMid := middleware.Authenticate(userService)
	app := fiber.New()
	app.Use(logger.New())

	api.NewAuth(app, userService, authMid)
	api.NewTransfer(app, transactionService, authMid)
	api.NewNotification(app, authMid, notificationService)
	api.NewTopUp(app, authMid, topupService)
	api.NewMidtrans(app, midtransService, topupService)

	sse.NewNotification(app, authMid, hub)

	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
