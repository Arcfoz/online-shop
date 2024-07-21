package product

import (
	"online-shop/infra/response"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        int       `db:"id"`
	SKU       string    `db:"sku"`
	Name      string    `db:"name"`
	Stock     int16     `db:"stock"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdataeAt time.Time `db:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func NewProductPaginationFromListProductRequest(req ListProductsRequestPayload) ProductPagination {
	req = req.GenerateDefaultValue()
	return ProductPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func (p Product) Validate() (err error) {
	if err = p.ValidateProductName(); err != nil {
		return
	}
	if err = p.ValidateProductStock(); err != nil {
		return
	}
	if err = p.ValidateProductPrice(); err != nil {
		return
	}
	return
}

func NewProductFromCreateProductRequest(req CreateProductRequestPayload) Product {
	return Product{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Stock:     req.Stock,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdataeAt: time.Now(),
	}
}

func (p Product) ValidateProductName() (err error) {
	if p.Name == "" {
		return response.ErrProductNameRequired
	}
	if len(p.Name) < 3 {
		return response.ErrProductNameInvalid
	}

	return
}

func (p Product) ValidateProductStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrProductStockInvalid
	}

	return
}

func (p Product) ValidateProductPrice() (err error) {
	if p.Price <= 0 {
		return response.ErrProductPriceInvalid
	}

	return
}
