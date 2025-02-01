package main

import (
	"os"

	"github.com/fakhrizalmus/tabungango/config"
	"github.com/fakhrizalmus/tabungango/controllers"
	"github.com/fakhrizalmus/tabungango/initializers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	initializers.LoadEnvVariables()
	Init()
}

func Init() {
	config.ConnectDatabase()
	app := fiber.New()

	app.Post("/daftar", controllers.Register)
	app.Post("/tabung", controllers.Tabung)
	app.Post("/tarik", controllers.Tarik)
	app.Get("/saldo/:no_rekening", controllers.Saldo)

	app.Listen(os.Getenv("PORT"))
}
