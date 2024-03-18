package app

import (
	"fmt"

	"github.com/b3nhard/ecommerce-api/internal/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	App    *fiber.App
	DB     *mongo.Client
	Config *config.Config
}

func NewApp(_db *mongo.Client, c *config.Config) *App {
	app := fiber.New(fiber.Config{
		AppName:       c.Server.AppName,
		CaseSensitive: true,
	})
	return &App{
		App: app,
		DB:  _db,
	}
}

func (a *App) Run() {
	a.App.Listen(fmt.Sprintf(":%d", a.Config.Server.Port))
}
