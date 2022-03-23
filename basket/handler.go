package basket

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	repo handlerRepo
}

type handlerRepo interface {
	AddProductToBasket(product_id int) error
	RemoveProductFromBasket(product_id int) error
	GetAllProducts() ([]Product, error)
	GetBasket() ([]ProductWrapper, error)
	ClearBasket() error
}

func NewHandler(_repo handlerRepo) *handler {
	return &handler{
		repo: _repo,
	}
}

func (h *handler) RegisterRoutes(app *fiber.App) {

	app.Get("api/products", h.GetAllProducts)
	app.Post("api/basket/:id", h.AddProductToBasket)
	app.Delete("api/basket/:id", h.RemoveProductFromBasket)
	app.Get("api/basket", h.GetBasket)
	app.Delete("api/basket", h.ClearBasket)

}

func (h *handler) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (h *handler) AddProductToBasket(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err := h.repo.AddProductToBasket(id)
	if err != nil {
		return ctx.SendStatus(400)
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handler) RemoveProductFromBasket(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	err = h.repo.RemoveProductFromBasket(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *handler) GetBasket(ctx *fiber.Ctx) error {
	basket, err := h.repo.GetBasket()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(basket)
}

func (h *handler) ClearBasket(ctx *fiber.Ctx) error {
	err := h.repo.ClearBasket()
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}
