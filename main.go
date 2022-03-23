package main

import (
	"basicShoppingCartApi/basket"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(port int) {
	app := fiber.New()
	app.Use(cors.New())
	repo := basket.NewRepo()
	handler := basket.NewHandler(repo)
	handler.RegisterRoutes(app)
	app.Listen(fmt.Sprintf(":%d", port))
}

func main() {
	StartServer(8080)
}
