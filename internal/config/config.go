package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("env")
	v.AllowEmptyEnv(false)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{
		Network: "tcp",
		Addr:    ":8080",
		Name:    "gendata-server",
	}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

type Config struct {
	RateLimit RateLimit `mapstructure:",squash"`
	TLS       TLS       `mapstructure:",squash"`
	Tracing   Tracing   `mapstructure:",squash"`
	Metrics   Metrics   `mapstructure:",squash"`
	Logger    Logger    `mapstructure:",squash"`
	Docs      Docs      `mapstructure:",squash"`
	// HTTP and GRPC server timeout [default = 30s]
	Timeout *time.Duration `mapstructure:"TIMEOUT"`
	Network string         `mapstructure:"NETWORK"`
	Addr    string         `mapstructure:"ADDR"`
	Debug   bool           `mapstructure:"DEBUG"`
	Name    string         `mapstructure:"-"`
	Version string         `mapstructure:"-"`
}

type TLS struct {
	Enabled bool `mapstructure:"TLS"`
	// Path to cert file
	Crt string `mapstructure:"TLS_CRT"`
	// Path to key file
	Key string `mapstructure:"TLS_KEY"`
}

type RateLimit struct {
	Disabled bool `mapstructure:"RATE_LIMIT_DISABLED"`
	// Window - defines time duration per window [default = 10s]
	Window *time.Duration `mapstructure:"RATE_LIMIT_WINDOW"`
	// Bucket - defines bucket number for each window [default = 100]
	Bucket *int32 `mapstructure:"RATE_LIMIT_BUCKET"`
	// CPUThreshold [default = 800]
	CPUThreshold *int64 `mapstructure:"RATE_LIMIT_CPU_THRESHOLD"`
}

type Logger struct {
	Disabled bool `mapstructure:"LOGGER_DISABLED"`
}

type Metrics struct {
	Disabled bool `mapstructure:"METRICS_DISABLED"`
}

type Docs struct {
	Disabled bool `mapstructure:"DOCS_DISABLED"`
}

type Tracing struct {
	Disabled bool `mapstructure:"TRACING_DISABLED"`
}
