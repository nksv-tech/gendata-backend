package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cfgPath = "../../.env.dist"
)

func TestLoadConfig(t *testing.T) {
	t.Run("filled config file", func(t *testing.T) {
		cfg, err := LoadConfig(cfgPath)
		assert.NoError(t, err)

		assert.Equal(t, cfg.Addr, ":8080")
		assert.Equal(t, cfg.Network, "tcp")
	})
	t.Run("empty config file (check default)", func(t *testing.T) {
		path := os.TempDir() + string(os.PathSeparator) + ".env"
		if err := os.WriteFile(path, []byte{}, 0666); err != nil {
			t.Fatal(err)
		}

		cfg, err := LoadConfig(path)
		assert.NoError(t, err)

		assert.Equal(t, cfg.Addr, ":8080")
		assert.Equal(t, cfg.Network, "tcp")
	})
}
