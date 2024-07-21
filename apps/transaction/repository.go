package transaction

import (
	"context"
	"database/sql"
	"online-shop/infra/response"

	"github.com/jmoiron/sqlx"
)

func newRepository(db *sqlx.DB) repository {
	repo := repository{
		db: db,
	}
	repo.createTableIfNotExists()
	return repo
}

func (r *repository) createTableIfNotExists() error {
	query := `
    CREATE TABLE IF NOT EXISTS transaction (
        id SERIAL PRIMARY KEY,
        user_public_id VARCHAR(100) NOT NULL,
        product_id INT NOT NULL,
        product_price INT NOT NULL,
        amount INT NOT NULL,
        subtotal INT NOT NULL,
        platform_fee INT NOT NULL,
        grand_total INT NOT NULL,
        status VARCHAR(50) NOT NULL,
        product_snapshot JSONB,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    )
    `
	_, err := r.db.Exec(query)
	return err
}

type repository struct {
	db *sqlx.DB
}

// GetTransactionByUserPublicId implements Repository.
func (r repository) GetTransactionByUserPublicId(ctx context.Context, userPublicId string) (trxs []Transaction, err error) {
	query := `
		SELECT
			id, user_public_id, product_id, product_price, amount, subtotal, platform_fee, grand_total, status, product_snapshot, created_at, updated_at
		FROM transaction
		WHERE user_public_id=$1
	`

	err = r.db.SelectContext(ctx, &trxs, query, userPublicId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrorNotFound
			return
		}
		return
	}

	return
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})
	return
}
func (repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Commit()
}
func (repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error) {
	query := `
		INSERT INTO transaction (
			user_public_id, product_id, product_price, amount, subtotal, platform_fee, grand_total, status, product_snapshot, created_at, updated_at
		) VALUES (
		 	:user_public_id, :product_id, :product_price, :amount, :subtotal, :platform_fee, :grand_total, :status, :product_snapshot, :created_at, :updated_at
		 )
	`
	stmt, err := tx.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, trx); err != nil {
		return
	}

	return
}

func (r repository) GetProductBySKU(ctx context.Context, productSKU string) (product Product, err error) {
	query := `
	SELECT
		id, sku, stock, price
	FROM products
	WHERE sku=$1
	`
	err = r.db.GetContext(ctx, &product, query, productSKU)
	if err != nil {
		if err == sql.ErrNoRows {
			return Product{}, response.ErrorNotFound
		}
		return
	}
	return
}

func (r repository) UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error) {
	query := `
	UPDATE products
	SET stock=:stock
	WHERE id=:id
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, product); err != nil {
		return
	}
	return
}
