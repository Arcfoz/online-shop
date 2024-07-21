package transaction

type CreateTransactionRequestPayload struct {
	ProductSKU   string `json:"product_sku"`
	UserPublicId string `json:"-"`
	Amount       uint8  `json:"amount"`
}
