package server

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/gen-data/gendata-server/internal/config"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	mwmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	mwratelimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	mwtracing "github.com/go-kratos/kratos/v2/middleware/tracing"
	mwvalidate "github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/soheilhy/cmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	defaultTimeout = time.Second * 30

	defaultRateLimiterWindow       = time.Second * 10
	defaultRateLimiterBucket       = int32(100)
	defaultRateLimiterCPUThreshold = int64(800)
)

var (
	// ProviderSet is server providers.
	ProviderSet = wire.NewSet(
		NewRateLimiter,
		NewMux,
		NewTLS,
		NewMiddlewares,
		NewGRPCServer,
		NewHTTPServer,
	)

	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duration(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

func init() {
	prometheus.MustRegister(_metricRequests, _metricSeconds)
}

func NewTLS(c *config.Config) *tls.Config {
	if !c.TLS.Enabled {
		return nil
	}
	cert, err := tls.LoadX509KeyPair(c.TLS.Crt, c.TLS.Key)
	if err != nil {
		panic(errors.WithMessage(err, "error TLS config"))
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
}

func NewMux(c *config.Config) cmux.CMux {
	lis, err := net.Listen(c.Network, c.Addr)
	if err != nil {
		panic(errors.WithMessagef(err, "error open listener on %s://%s", c.Network, c.Addr))
	}
	mux := cmux.New(lis)
	go func() {
		err := mux.Serve()
		if err != nil {
			panic(err)
		}
	}()
	return mux
}

func NewRateLimiter(c *config.Config) ratelimit.Limiter {
	if c.RateLimit.Disabled {
		return nil
	}

	window := defaultRateLimiterWindow
	if c.RateLimit.Window != nil {
		window = *c.RateLimit.Window
	}

	bucket := defaultRateLimiterBucket
	if c.RateLimit.Bucket != nil {
		bucket = *c.RateLimit.Bucket
	}

	cpuThreshold := defaultRateLimiterCPUThreshold
	if c.RateLimit.CPUThreshold != nil {
		cpuThreshold = *c.RateLimit.CPUThreshold
	}

	return bbr.NewLimiter(
		bbr.WithWindow(window),
		bbr.WithBucket(int(bucket)),
		bbr.WithCPUThreshold(cpuThreshold),
	)
}

func NewMiddlewares(
	c *config.Config,
	logger log.Logger,
	limiter ratelimit.Limiter,
) []middleware.Middleware {
	mws := []middleware.Middleware{
		recovery.Recovery(),
	}
	if !c.Tracing.Disabled {
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(c.Name),
			)),
		)
		otel.SetTracerProvider(tp)
		mws = append(mws, mwtracing.Server())
	}
	if !c.Logger.Disabled && logger != nil {
		mws = append(mws, logging.Server(logger))
	}
	if !c.RateLimit.Disabled && limiter != nil {
		mws = append(mws, mwratelimit.Server(mwratelimit.WithLimiter(limiter)))
	}
	if !c.Metrics.Disabled {
		mws = append(mws, mwmetrics.Server(
			mwmetrics.WithRequests(prom.NewCounter(_metricRequests)),
			mwmetrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
		))
	}

	mws = append(mws, mwvalidate.Validator())

	return mws
}
