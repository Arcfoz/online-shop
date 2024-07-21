package product

import (
	"online-shop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:  "White Tshirt",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("product name required", func(t *testing.T) {
		product := Product{
			Name:  "",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameRequired, err)
	})

	t.Run("product name invalid", func(t *testing.T) {
		product := Product{
			Name:  "ab",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameInvalid, err)
	})

	t.Run("product stock invalid", func(t *testing.T) {
		product := Product{
			Name:  "White Tshirt",
			Stock: 0,
			Price: 100_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductStockInvalid, err)
	})

	t.Run("product price invalid", func(t *testing.T) {
		product := Product{
			Name:  "White Tshirt",
			Stock: 10,
			Price: 0,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductPriceInvalid, err)
	})

}
