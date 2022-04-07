package main

import (
	"flag"
	"os"

	"github.com/gen-data/gendata-server/internal/config"
	kzap "github.com/gen-data/gendata-server/internal/utils/zap"
	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
// go build -ldflags "-X main.Name=gen-data"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", ".env", "config path, eg: -conf .env")
}

func newApp(logger log.Logger, g *grpc.Server, h *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(g, h),
	)
}

func main() {
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(flagconf)
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg := new(config.Config)
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	var zapL *zap.Logger
	{
		zapCfg := zap.NewProductionConfig()
		if cfg.Debug {
			zapCfg = zap.NewDevelopmentConfig()
		}
		var err error
		zapL, err = zapCfg.Build()
		if err != nil {
			panic(err)
		}
	}

	kLog := kzap.NewLogger(zapL)
	logger := log.With(kLog,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
		)),
	)
	otel.SetTracerProvider(tp)

	app, cleanup, err := wireApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
