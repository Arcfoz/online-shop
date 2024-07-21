package auth

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
	repo.createTableIfNotExist()
	return repo
}

func (r repository) createTableIfNotExist() error {
	query := `
        CREATE TABLE IF NOT EXISTS auth (
            id SERIAL PRIMARY KEY,
            public_id VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(50) NOT NULL,
            created_at TIMESTAMP NOT NULL,
            updated_at TIMESTAMP NOT NULL
        )
    `
	_, err := r.db.Exec(query)
	return err
}

func (r repository) CreateAuth(ctx context.Context, model AuthEntity) (err error) {
	query := `
		INSERT INTO auth (
			public_id, email, password, role, created_at, updated_at
		) VALUES (
		 	:public_id, :email, :password, :role, :created_at, :updated_at
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

// GetAuthByEmail implements Repository
func (r repository) GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error) {
	query := `
		SELECT
			id, public_id, email, password, role, created_at, updated_at
		FROM auth
		WHERE email=$1
	`
	err = r.db.GetContext(ctx, &model, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}
