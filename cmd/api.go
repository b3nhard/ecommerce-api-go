package main

import (
	"fmt"
	"os"

	"github.com/b3nhard/ecommerce-api/internal/app"
	"github.com/b3nhard/ecommerce-api/internal/config"
	"github.com/b3nhard/ecommerce-api/internal/database"
	"github.com/b3nhard/ecommerce-api/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	c := config.NewConfig()
	db := database.InitDb()
	app := app.NewApp(db, c)

	// Setup router
	router.SetupRouter(app, db)

	// Run Application
	app.App.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
