package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/ewallet/internal/api"
	"github.com/khairulharu/ewallet/internal/component"
	"github.com/khairulharu/ewallet/internal/config"
	"github.com/khairulharu/ewallet/internal/middleware"
	"github.com/khairulharu/ewallet/internal/repository"
	"github.com/khairulharu/ewallet/internal/service"
)

func main() {
	config := config.Get()
	dbConnection := component.GetDatabaseConnection(config)
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRepository, cacheConnection)

	authMid := middleware.Authenticate(userService)
	app := fiber.New()
	app.Use(logger.New())

	api.NewAuth(app, userService, authMid)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
