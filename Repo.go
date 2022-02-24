package main

import (
	models "basicShoppingCartApi/models"
	"errors"
	"sort"
)

var products models.Products
var basket *models.Basket

type Repo struct{}

func InitRepo(repo *Repo) {
	products = make(models.Products)
	models.PopulateProducts(products)
	basket = new(models.Basket)
	basket.ProductWrappers = make(map[int]models.ProductWrapper)
}

func (repo *Repo) AddProductToBasket(product_id int) error {

	if _, exists := products[product_id]; !exists {
		return errors.New("no product with given id")
	}
	if productWrapper, exists := basket.ProductWrappers[product_id]; exists {
		productWrapper.ProductCount++
		basket.ProductWrappers[product_id] = productWrapper
		return nil
	}
	basket.ProductWrappers[product_id] = models.ProductWrapper{
		Product:      products[product_id],
		ProductCount: 1,
	}
	return nil
}

func (repo *Repo) RemoveProductFromBasket(product_id int) error {
	if _, exists := products[product_id]; !exists {
		return errors.New("no product with given id")
	}
	if productWrapper, exists := basket.ProductWrappers[product_id]; exists {
		if productWrapper.ProductCount == 1 {
			delete(basket.ProductWrappers, product_id)
			return nil
		}
		productWrapper.ProductCount--
		basket.ProductWrappers[product_id] = productWrapper
		return nil
	} else {
		return errors.New("no product in basket with given id")
	}
}

func (repo *Repo) GetAllProducts() []models.Product {
	allProducts := make([]models.Product, 0, len(products))
	for i := 1; i < len(products)+1; i++ {
		allProducts = append(allProducts, products[i])
	}
	return allProducts
}

func (repo *Repo) GetBasket() []models.ProductWrapper {
	keys := make([]int, 0, len(basket.ProductWrappers))
	for k := range basket.ProductWrappers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var productWrappers []models.ProductWrapper

	for _, k := range keys {
		productWrappers = append(productWrappers, basket.ProductWrappers[k])
	}
	return productWrappers
}

func (repo *Repo) ClearBasket() {
	for key := range basket.ProductWrappers {
		delete(basket.ProductWrappers, key)
	}
}
