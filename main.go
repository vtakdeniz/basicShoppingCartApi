package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func createApp(repo *Repo) *fiber.App {
	InitRepo(repo)
	app := fiber.New()
	app.Use(cors.New())
	app.Get("api/products", func(c *fiber.Ctx) error {
		allProducts := repo.GetAllProducts()
		return c.JSON(allProducts)
	})
	app.Get("api/basket/add/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.New("error reading url parameter")
		}
		err = repo.AddProductToBasket(id)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendStatus(fiber.StatusOK)
	})
	app.Get("api/basket/remove/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.New("error reading url parameter")
		}
		err = repo.RemoveProductFromBasket(id)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("api/basket", func(c *fiber.Ctx) error {
		return c.JSON(repo.GetBasket())
	})
	app.Get("api/basket/clear", func(c *fiber.Ctx) error {
		repo.ClearBasket()
		return c.JSON(fiber.StatusOK)
	})
	return app
}

func main() {
	app := createApp(new(Repo))
	app.Listen(":8080")
}
