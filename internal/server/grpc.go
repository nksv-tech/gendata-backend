package server

import (
	"crypto/tls"
	"math"
	"time"

	"github.com/gen-data/gendata-server/internal/config"
	"github.com/gen-data/gendata-server/internal/service"
	v1 "github.com/gen-data/gendata-server/pkg/gendata/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	tgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	maxRecvMsgSize = math.MaxInt32
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *config.Config,
	srv *service.GenDataService,
	mws []middleware.Middleware,
	mux cmux.CMux,
	logger log.Logger,
	tlsCfg *tls.Config,
) *tgrpc.Server {
	var opts = []tgrpc.ServerOption{
		tgrpc.Listener(
			mux.MatchWithWriters(
				cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
			),
		),
		tgrpc.Middleware(mws...),
		tgrpc.Options(
			grpc.MaxRecvMsgSize(maxRecvMsgSize),
		),
		tgrpc.Logger(logger),
		tgrpc.TLSConfig(tlsCfg),
	}

	timeout := time.Second * 30
	if c.Timeout != nil {
		timeout = *c.Timeout
	}
	opts = append(opts, tgrpc.Timeout(timeout))

	serv := tgrpc.NewServer(opts...)
	v1.RegisterGenDataServiceServer(serv, srv)
	return serv
}
