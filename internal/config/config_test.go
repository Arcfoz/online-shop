package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		fillename := "../../cmd/api/config.yaml"
		err := LoadConfig(fillename)

		require.Nil(t, err)
		log.Printf("%+v\n", Cfg)
	})

	t.Run("File not exist", func(t *testing.T) {
		fillename := "config.yaml"
		err := LoadConfig(fillename)

		require.NotNil(t, err)
		log.Printf("%+v\n", Cfg)
	})
}