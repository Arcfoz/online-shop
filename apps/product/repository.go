package product

import (
	"context"
	"database/sql"
	"online-shop/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	repo := repository{
		db: db,
	}
	repo.createTableIfNotExists()
	return repo
}

func (r *repository) createTableIfNotExists() error {
	query := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        sku VARCHAR(255) UNIQUE NOT NULL,
        name VARCHAR(255) NOT NULL,
        stock INT NOT NULL,
        price INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    )
    `
	_, err := r.db.Exec(query)
	return err
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			sku, name, stock, price,created_at, updated_at
		) VALUES (
		  :sku, :name, :stock, :price,:created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := `
		SELECT
			id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE id>$1
		ORDER BY id ASC
		LIMIT $2
	`
	err = r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return
	}
	return
}

func (r repository) GetProductBySKU(ctx context.Context, sku string) (product Product, err error) {
	query := `
		SELECT
			id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, sku)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}
