package main

import (
	"flag"
	"os"

	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/nksv-tech/gendata-backend/internal/config"
	kzap "github.com/nksv-tech/gendata-backend/internal/utils/zap"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
// go build -ldflags "-X main.Name=gen-data"
// go build -ldflags "-X main.Commit=hash"
// go build -ldflags "-X main.Date=2022-04-21"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// Commit is the git commit of the compiled software.
	Commit string
	// Date is the date of the compiled software.
	Date string
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

	cfg, err := config.LoadConfig(flagconf)
	if err != nil {
		panic(err)
	}
	cfg.Name = Name
	cfg.Version = Version

	var zapL *zap.Logger
	{
		zapCfg := zap.NewProductionConfig()
		if cfg.Debug {
			zapCfg = zap.NewDevelopmentConfig()
		}
		zapL, err = zapCfg.Build()
		if err != nil {
			panic(err)
		}
	}

	kLog := kzap.NewLogger(zapL)
	defer func(kLog *kzap.Logger) {
		if err = kLog.Sync(); err != nil {
			panic(err)
		}
	}(kLog)

	logger := log.With(kLog,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"service.commit", Commit,
		"service.build_date", Date,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
