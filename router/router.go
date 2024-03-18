package router

import (
	"github.com/b3nhard/ecommerce-api/internal/app"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(app *app.App, db *mongo.Client) {
	app.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Ecommerce API",
		})
	})
}
