package product

import "time"

type ProductsListResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductsListResponseFromEntity(products []Product) []ProductsListResponse {
	productsList := []ProductsListResponse{}

	for _, product := range products {
		productsList = append(productsList, ProductsListResponse{
			Id:    product.Id,
			SKU:   product.SKU,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}

	return productsList
}

type ProductDetailResponse struct {
	Id        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Stock     int16     `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdataeAt time.Time `json:"updated_at"`
}
