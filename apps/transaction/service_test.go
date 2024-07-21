package transaction

import (
	"context"
	"online-shop/external/database"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "8c3c6d05-422a-4622-a46d-f9a5dad02b73",
			Amount:       2,
			UserPublicId: "5c156934-5d52-465c-83fc-ee4ff6053700",
		}
		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)

	})
}
