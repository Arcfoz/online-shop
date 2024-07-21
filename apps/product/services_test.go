package product

import (
	"context"
	"log"
	"online-shop/external/database"
	"online-shop/infra/response"
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

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Abibas Airrobic",
		Stock: 10,
		Price: 300_000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Failed(t *testing.T) {
	t.Run("product name is required", func(t *testing.T) {
		name := ""
		req := CreateProductRequestPayload{
			Name:  name,
			Stock: 10,
			Price: 300_000,
		}
		err := svc.CreateProduct(context.Background(), req)

		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameRequired, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductsRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)

	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestProductDetail_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "(test)",
		Stock: 10,
		Price: 300_000,
	}

	ctx := context.Background()
	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	products, err := svc.ListProducts(ctx, ListProductsRequestPayload{
		Cursor: 0,
		Size:   10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	
	require.Nil(t, err)
	require.NotEmpty(t, product)

	log.Printf("%+v", product)
}
