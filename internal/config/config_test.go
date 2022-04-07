package config

import (
	"context"
	"encoding/json"
	"testing"

	v1 "github.com/gen-data/gendata-server/pkg/gendata/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const (
	cfgPath = "../../.env.dist"
)

func TestLoadConfig(t *testing.T) {
	exp := `{"TLS":{"Enabled":false,"Crt":"","Key":""},"Metrics":{"Disabled":false},"Logger":{"Disabled":false},"RateLimit":{"Disabled":true,"Window":10000000000,"Bucket":100,"CPUThreshold":800},"Docs":{"Disabled":false},"Timeout":30000000000,"Network":"tcp","Addr":":8080","Debug":true}`
	cfg, err := LoadConfig(cfgPath)
	assert.NoError(t, err)
	bs, err := json.Marshal(cfg)
	assert.JSONEq(t, exp, string(bs))
}

func TestAa(t *testing.T) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cl := v1.NewGenDataServiceClient(conn)
	_, err = cl.PredefinedSettings(context.Background(), &v1.PredefinedSettingsRequest{})
	if err != nil {
		panic(err)
	}
}
