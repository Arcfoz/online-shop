package transaction

import (
	"encoding/json"
	"online-shop/infra/response"
	"time"
)

type TransactionStatus uint8

const (
	TransactionStatus_Created    TransactionStatus = 1
	TransactionStatus_Progress   TransactionStatus = 10
	TransactionStatus_InDelivery TransactionStatus = 15
	TransactionStatus_Completed  TransactionStatus = 20

	TRX_CREATED    string = "CREATED"
	TRX_PROGRESS   string = "ON PROGRESS"
	TRX_INDELIVERY string = "IN DELIVERY"
	TRX_COMPLETED  string = "COMPLETE"
	TRX_UNKNOWN    string = "UNKNOWN STATUS"
)

var (
	MappingTransactionStatus = map[TransactionStatus]string{
		TransactionStatus_Created:    TRX_CREATED,
		TransactionStatus_Progress:   TRX_PROGRESS,
		TransactionStatus_InDelivery: TRX_INDELIVERY,
		TransactionStatus_Completed:  TRX_COMPLETED,
	}
)

type Transaction struct {
	Id           int               `db:"id"`
	UserPublicId string            `db:"user_public_id"`
	ProductId    uint              `db:"product_id"`
	ProductPrice uint              `db:"product_price"`
	Amount       uint8             `db:"amount"`
	SubTotal     uint              `db:"subtotal"`
	PlatformFee  uint              `db:"platform_fee"`
	GrandTotal   uint              `db:"grand_total"`
	Status       TransactionStatus `db:"status"`
	ProductJSON  json.RawMessage   `db:"product_snapshot"`
	CreatedAt    time.Time         `db:"created_at"`
	UpdataeAt    time.Time         `db:"updated_at"`
}

func NewTransaction(userPublicId string) Transaction {
	return Transaction{
		UserPublicId: userPublicId,
		Status:       TransactionStatus_Created,
		CreatedAt:    time.Now(),
		UpdataeAt:    time.Now(),
	}
}

func NewTransactionFromCreateRequest(req CreateTransactionRequestPayload) Transaction {
	return Transaction{
		UserPublicId: req.UserPublicId,
		Amount:       req.Amount,
		Status:       TransactionStatus_Created,
		CreatedAt:    time.Now(),
		UpdataeAt:    time.Now(),
	}
}

func (t *Transaction) SetPlatformFee(platformFee uint) *Transaction {
	t.PlatformFee = platformFee
	return t
}

func (t Transaction) Validate() (err error) {
	if t.Amount == 0 {
		return response.ErrorInvalidAmount
	}

	return
}

func (t Transaction) ValidateStock(productStock uint8) (err error) {
	if t.Amount > productStock {
		return response.ErrAmountGreaterThanStock
	}

	return
}

func (t *Transaction) SetSubTotal() {
	if t.SubTotal == 0 {
		t.SubTotal = t.ProductPrice * uint(t.Amount)
	}
}

// set subtotal and grand total
func (t *Transaction) SetGrandTotal() {
	if t.GrandTotal == 0 {
		t.SetSubTotal()
		t.GrandTotal = t.SubTotal + t.PlatformFee
	}

}

// set product id, price, and json
func (t *Transaction) FromProduct(product Product) *Transaction {
	t.ProductId = uint(product.Id)
	t.ProductPrice = uint(product.Price)

	t.SetProductJSON(product)
	return t
}

func (t *Transaction) SetProductJSON(product Product) (err error) {
	productJson, err := json.Marshal(product)
	if err != nil {
		return
	}

	t.ProductJSON = productJson

	return
}

func (t Transaction) GetProduct() (product Product, err error) {
	err = json.Unmarshal(t.ProductJSON, &product)
	if err != nil {
		return
	}
	return
}

func (t Transaction) GetStatus() string {
	status, ok := MappingTransactionStatus[t.Status]
	if !ok {
		return TRX_UNKNOWN
	}
	return status
}

func (t Transaction) ToTransactionHistoryResponse() TransactionHistoryResponse {
	product, err := t.GetProduct()
	if err != nil {
		product = Product{}
	}
	return TransactionHistoryResponse{
		Id:           t.Id,
		UserPublicId: t.UserPublicId,
		ProductId:    t.ProductId,
		ProductPrice: t.ProductPrice,
		Amount:       t.Amount,
		SubTotal:     t.SubTotal,
		PlatformFee:  t.PlatformFee,
		GrandTotal:   t.GrandTotal,
		Status:       t.GetStatus(),
		CreatedAt:    t.CreatedAt,
		UpdataeAt:    t.UpdataeAt,
		Product:      product,
	}
}
