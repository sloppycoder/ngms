package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sloppycoder/ngms/accounts/app/server"
	"net/http"
	"os"

	"google.golang.org/grpc/grpclog"
)

func main() {
	err := run()
	if err != nil {
		grpclog.Errorf("server was shutdown with errors: %v", err)
		os.Exit(1)
	}
}

func prometheusMiddleWare(next http.Handler) http.Handler {
	handler := promhttp.Handler()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			handler.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func run() error {
	// Application context
	ctx := context.Background()

	// Create some standard server metrics.
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcAddr("tcp", ":3001"),
		grapiserver.WithGatewayAddr("tcp", ":3000"),
		grapiserver.WithGatewayServerMiddlewares(prometheusMiddleWare),
		grapiserver.WithGrpcServerStreamInterceptors(grpcMetrics.StreamServerInterceptor()),
		grapiserver.WithGrpcServerUnaryInterceptors(grpcMetrics.UnaryServerInterceptor()),
		grapiserver.WithServers(
			server.NewAccountServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
