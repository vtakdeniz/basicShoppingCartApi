package basket

import (
	"errors"
	"sort"
)

type Repository struct {
	Products Products
	Basket   Basket
}

func NewRepo() *Repository {
	products := make(Products)
	PopulateProducts(products)
	basket := new(Basket)
	basket.ProductWrappers = make(map[int]ProductWrapper)
	return &Repository{
		Products: products,
		Basket:   *basket,
	}
}

func (repo *Repository) AddProductToBasket(product_id int) error {

	if _, exists := repo.Products[product_id]; !exists {
		return errors.New("no product with given id")
	}
	if productWrapper, exists := repo.Basket.ProductWrappers[product_id]; exists {
		productWrapper.ProductCount++
		repo.Basket.ProductWrappers[product_id] = productWrapper
		return nil
	}
	repo.Basket.ProductWrappers[product_id] = ProductWrapper{
		Product:      repo.Products[product_id],
		ProductCount: 1,
	}
	return nil
}

func (repo *Repository) RemoveProductFromBasket(product_id int) error {
	if _, exists := repo.Products[product_id]; !exists {
		return errors.New("no product with given id")
	}
	if productWrapper, exists := repo.Basket.ProductWrappers[product_id]; exists {
		if productWrapper.ProductCount == 1 {
			delete(repo.Basket.ProductWrappers, product_id)
			return nil
		}
		productWrapper.ProductCount--
		repo.Basket.ProductWrappers[product_id] = productWrapper
		return nil
	} else {
		return errors.New("no product in basket with given id")
	}
}

func (repo *Repository) GetAllProducts() ([]Product, error) {
	allProducts := make([]Product, 0, len(repo.Products))
	for i := 1; i < len(repo.Products)+1; i++ {
		allProducts = append(allProducts, repo.Products[i])
	}
	return allProducts, nil
}

func (repo *Repository) GetBasket() ([]ProductWrapper, error) {
	keys := make([]int, 0, len(repo.Basket.ProductWrappers))
	for k := range repo.Basket.ProductWrappers {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var productWrappers []ProductWrapper

	for _, k := range keys {
		productWrappers = append(productWrappers, repo.Basket.ProductWrappers[k])
	}
	return productWrappers, nil
}

func (repo *Repository) ClearBasket() error {
	for key := range repo.Basket.ProductWrappers {
		delete(repo.Basket.ProductWrappers, key)
	}
	return nil
}
