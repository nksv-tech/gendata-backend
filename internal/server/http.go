package server

import (
	"crypto/tls"
	"embed"
	"io"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/gen-data/gendata-server/internal/config"
	"github.com/gen-data/gendata-server/internal/service"
	v1 "github.com/gen-data/gendata-server/pkg/gendata/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
)

//go:generate cp -r ../../docs ./docs
//go:embed *
var docs embed.FS

// NewHTTPServer new as HTTP server.
func NewHTTPServer(
	c *config.Config,
	srv *service.GenDataService,
	mws []middleware.Middleware,
	mux cmux.CMux,
	logger log.Logger,
	tlsCfg *tls.Config,
) *thttp.Server {
	var opts = []thttp.ServerOption{
		thttp.Listener(mux.Match(cmux.HTTP1Fast())),
		thttp.Middleware(mws...),
		thttp.Logger(logger),
		thttp.TLSConfig(tlsCfg),
	}

	timeout := defaultTimeout
	if c.Timeout != nil {
		timeout = *c.Timeout
	}
	opts = append(opts, thttp.Timeout(timeout))

	s := thttp.NewServer(opts...)
	v1.RegisterGenDataServiceHTTPServer(s, srv)

	if !c.Metrics.Disabled {
		s.Handle("/metrics", promhttp.Handler())
	}

	if c.Debug {
		s.HandleFunc("/debug/pprof", pprof.Index)
		s.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		s.HandleFunc("/debug/pprof/profile", pprof.Profile)
		s.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		s.HandleFunc("/debug/pprof/trace", pprof.Trace)
		s.HandleFunc("/debug/allocs", pprof.Handler("allocs").ServeHTTP)
		s.HandleFunc("/debug/block", pprof.Handler("block").ServeHTTP)
		s.HandleFunc("/debug/goroutine", pprof.Handler("goroutine").ServeHTTP)
		s.HandleFunc("/debug/heap", pprof.Handler("heap").ServeHTTP)
		s.HandleFunc("/debug/mutex", pprof.Handler("mutex").ServeHTTP)
		s.HandleFunc("/debug/threadcreate", pprof.Handler("threadcreate").ServeHTTP)

		s.HandleFunc("/debug/logs", func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.Copy(w, os.Stdout)
		})
	}

	if !c.Docs.Disabled {
		s.Handle("/docs/openapi.yaml", http.FileServer(http.FS(docs)))
	}

	return s
}
