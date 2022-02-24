package main

import (
	models "basicShoppingCartApi/models"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	resTestGetProducts             = `[{"Id":1,"category":"Home","title":"Wooden dinner table","description":"Wooden table for dinner room","price":1250,"pic_link":"https://ronixtools.com/en/blog/wp-content/uploads/2021/03/Learn-how-to-make-a-simple-wooden-table-at-home1.jpg"},{"Id":2,"category":"Garden","title":"Garden Hose","description":"Green garden hose","price":50,"pic_link":"https://cdn11.bigcommerce.com/s-3c8l9ljcjn/images/stencil/1280x1280/products/24929/35825/39a0105_lifetime_garden_hose_rack_alt2__22440__05859.1593445624__65789.1631907713.jpg?c=1"},{"Id":3,"category":"Garden","title":"Wooden garden table","description":"Wooden table for gardens","price":1350,"pic_link":"https://image.made-in-china.com/202f0j00pwvEBLQMaTqP/Outdoor-Furniture-Wooden-Garden-Table-Picnic-Table-Sets-for-Children.jpg"},{"Id":4,"category":"Home","title":"TV Unit","description":"TV Unit for plasma TV's","price":150,"pic_link":"https://www.ulcdn.net/opt/www.ulcdn.net/images/products/125614/slide/666x363/Zephyr_LargeTV_Unit_TK_2.jpg?1608823365"},{"Id":5,"category":"Cleaning","title":"Robot vacuum","description":"Robot vacuum that clean your house automatically","price":700,"pic_link":"https://images.hepsiburada.net/assets/ProductDescription/202007/52168f7d-a5c5-47c3-b810-3810d3b57e8e.jpg"},{"Id":6,"category":"Electronics","title":"Iphone 13","description":"Latest model of iphone released this year","price":1000,"pic_link":"https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-13-pro-silver-select?wid=470\u0026hei=556\u0026fmt=jpeg\u0026qlt=95\u0026.v=1631652954000"},{"Id":7,"category":"Electronics","title":"Tv","description":"Plasma Tv","price":1000,"pic_link":"https://5.imimg.com/data5/NA/OH/MY-2906751/td-500x500.jpg"},{"Id":8,"category":"Electronics","title":"Imac pro","description":"Latest release of Apple's Imac","price":3000,"pic_link":"https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/v2-89592_large.jpg"},{"Id":9,"category":"Home","title":"Rug","description":"Self cleaning rug for houses","price":1100,"pic_link":"https://www.therange.co.uk/_m5/5/9/1596549272_2_2067.jpg"}]`
	resTestAddToBasket             = `[{"product":{"Id":5,"category":"Cleaning","title":"Robot vacuum","description":"Robot vacuum that clean your house automatically","price":700,"pic_link":"https://images.hepsiburada.net/assets/ProductDescription/202007/52168f7d-a5c5-47c3-b810-3810d3b57e8e.jpg"},"productCount":2},{"product":{"Id":6,"category":"Electronics","title":"Iphone 13","description":"Latest model of iphone released this year","price":1000,"pic_link":"https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-13-pro-silver-select?wid=470\u0026hei=556\u0026fmt=jpeg\u0026qlt=95\u0026.v=1631652954000"},"productCount":1}]`
	resTestRemoveProductFromBasket = `[{"product":{"Id":4,"category":"Home","title":"TV Unit","description":"TV Unit for plasma TV's","price":150,"pic_link":"https://www.ulcdn.net/opt/www.ulcdn.net/images/products/125614/slide/666x363/Zephyr_LargeTV_Unit_TK_2.jpg?1608823365"},"productCount":1}]`
)

func TestGetProducts(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		body         string
	}{
		{
			description:  "Receivers all hardcoded products",
			route:        "/api/products",
			expectedCode: 200,
			body:         resTestGetProducts,
		},
	}

	repo := new(Repo)
	InitRepo(repo)
	app := createApp(repo)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		body, _ := ioutil.ReadAll(resp.Body)
		bodyObj := []models.Product{}

		json.Unmarshal(body, &bodyObj)
		bodyJSON, _ := json.Marshal(bodyObj)
		bodyJSONString := string(bodyJSON)

		assert.JSONEq(t, test.body, bodyJSONString, test.description)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestAddToBasket(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
	}{
		{
			description:  "Adds product to the basket correctly ",
			expectedCode: 200,
		},
	}

	repo := new(Repo)
	InitRepo(repo)
	app := createApp(repo)

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/api/basket/add/5", nil)
		app.Test(req, 1)
		req2 := httptest.NewRequest("GET", "/api/basket/add/5", nil)
		app.Test(req2, 1)
		req3 := httptest.NewRequest("GET", "/api/basket/add/6", nil)
		app.Test(req3, 1)

		reqBasket := httptest.NewRequest("GET", "/api/basket", nil)
		respBasket, _ := app.Test(reqBasket, 1)

		body, _ := ioutil.ReadAll(respBasket.Body)

		assert.JSONEq(t, resTestAddToBasket, string(body), test.description)
		assert.Equal(t, test.expectedCode, respBasket.StatusCode, test.description)
	}
}

func TestAddToBasketUnavailableProduct(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
	}{
		{
			description:  "Sends back bad request if product doesn't exists",
			expectedCode: 400,
		},
	}

	repo := new(Repo)
	InitRepo(repo)
	app := createApp(repo)

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/api/basket/add/5423", nil)
		resp, _ := app.Test(req, 1)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestRemoveProductFromBasket(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		body         string
	}{
		{
			description:  "Removes product from basket correctly",
			expectedCode: 200,
			body:         resTestRemoveProductFromBasket,
		},
	}

	repo := new(Repo)
	InitRepo(repo)
	app := createApp(repo)

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/api/basket/add/6", nil)
		app.Test(req, 1)
		req2 := httptest.NewRequest("GET", "/api/basket/add/4", nil)
		app.Test(req2, 1)
		req3 := httptest.NewRequest("GET", "/api/basket/add/4", nil)
		app.Test(req3, 1)
		req4 := httptest.NewRequest("GET", "/api/basket/remove/4", nil)
		app.Test(req4, 1)
		req5 := httptest.NewRequest("GET", "/api/basket/remove/6", nil)
		app.Test(req5, 1)
		reqFinal := httptest.NewRequest("GET", "/api/basket", nil)
		resp, _ := app.Test(reqFinal, 1)
		body, _ := ioutil.ReadAll(resp.Body)
		assert.JSONEq(t, test.body, string(body), test.description)
	}
}

func TestClearsBasket(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		body         string
	}{
		{
			description:  "Removes product from basket correctly",
			expectedCode: 200,
			body:         "null",
		},
	}

	repo := new(Repo)
	InitRepo(repo)
	app := createApp(repo)

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/api/basket/clear", nil)
		app.Test(req, 1)

		reqFinal := httptest.NewRequest("GET", "/api/basket", nil)
		resp, _ := app.Test(reqFinal, 1)

		body, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, test.body, string(body), test.description)
	}
}
