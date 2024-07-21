package database

import (
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	fillename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(fillename)

	if err != nil{
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T){
	t.Run("Success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

	// t.Run("Invalid password", func(t *testing.T) {
	// 	cfg := config.Cfg.DB
	// 	cfg.Password = "invalid password"
	// 	db, err := ConnectPostgres(cfg)
	// 	require.NotNil(t, err)
	// 	require.Nil(t, db)
	// })
}