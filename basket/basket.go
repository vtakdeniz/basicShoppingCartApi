package basket

type Product struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Pic_link    string `json:"pic_link"`
}

type Products map[int]Product

type ProductWrapper struct {
	Product      Product `json:"product"`
	ProductCount int     `json:"productCount"`
}

type Basket struct {
	ProductWrappers map[int]ProductWrapper
}

func PopulateProducts(products Products) {
	products[1] = Product{
		Id:          1,
		Category:    "Home",
		Title:       "Wooden dinner table",
		Description: "Wooden table for dinner room",
		Price:       1250,
		Pic_link:    "https://ronixtools.com/en/blog/wp-content/uploads/2021/03/Learn-how-to-make-a-simple-wooden-table-at-home1.jpg",
	}
	products[2] = Product{
		Id:          2,
		Category:    "Garden",
		Title:       "Garden Hose",
		Description: "Green garden hose",
		Price:       50,
		Pic_link:    "https://cdn11.bigcommerce.com/s-3c8l9ljcjn/images/stencil/1280x1280/products/24929/35825/39a0105_lifetime_garden_hose_rack_alt2__22440__05859.1593445624__65789.1631907713.jpg?c=1",
	}
	products[3] = Product{
		Id:          3,
		Category:    "Garden",
		Title:       "Wooden garden table",
		Description: "Wooden table for gardens",
		Price:       1350,
		Pic_link:    "https://image.made-in-china.com/202f0j00pwvEBLQMaTqP/Outdoor-Furniture-Wooden-Garden-Table-Picnic-Table-Sets-for-Children.jpg",
	}
	products[4] = Product{
		Id:          4,
		Category:    "Home",
		Title:       "TV Unit",
		Description: "TV Unit for plasma TV's",
		Price:       150,
		Pic_link:    "https://www.ulcdn.net/opt/www.ulcdn.net/images/products/125614/slide/666x363/Zephyr_LargeTV_Unit_TK_2.jpg?1608823365",
	}
	products[5] = Product{
		Id:          5,
		Category:    "Cleaning",
		Title:       "Robot vacuum",
		Description: "Robot vacuum that clean your house automatically",
		Price:       700,
		Pic_link:    "https://images.hepsiburada.net/assets/ProductDescription/202007/52168f7d-a5c5-47c3-b810-3810d3b57e8e.jpg",
	}
	products[6] = Product{
		Id:          6,
		Category:    "Electronics",
		Title:       "Iphone 13",
		Description: "Latest model of iphone released this year",
		Price:       1000,
		Pic_link:    "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-13-pro-silver-select?wid=470&hei=556&fmt=jpeg&qlt=95&.v=1631652954000",
	}
	products[7] = Product{
		Id:          7,
		Category:    "Electronics",
		Title:       "Tv",
		Description: "Plasma Tv",
		Price:       1000,
		Pic_link:    "https://5.imimg.com/data5/NA/OH/MY-2906751/td-500x500.jpg",
	}
	products[8] = Product{
		Id:          8,
		Category:    "Electronics",
		Title:       "Imac pro",
		Description: "Latest release of Apple's Imac",
		Price:       3000,
		Pic_link:    "https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/v2-89592_large.jpg",
	}
	products[9] = Product{
		Id:          9,
		Category:    "Home",
		Title:       "Rug",
		Description: "Self cleaning rug for houses",
		Price:       1100,
		Pic_link:    "https://www.therange.co.uk/_m5/5/9/1596549272_2_2067.jpg",
	}
}
