package transaction

import "online-shop/infra/response"

type Product struct {
	Id    int    `db:"id" json:"id"`
	SKU   string `db:"sku" json:"sku"`
	Name  string `db:"name" json:"name"`
	Stock uint   `db:"stock" json:"-"`
	Price int    `db:"price" json:"price"`
}

func (p Product) IsExist() bool {
	return p.Id != 0
}

func (p *Product) UpdateStockProduct(amount uint8) (err error) {
	if p.Stock <= uint(amount) {
		return response.ErrorAmountGreaterThanStock
	}
	p.Stock = p.Stock - uint(amount)
	return
}
