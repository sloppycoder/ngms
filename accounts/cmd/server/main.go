package main

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/sloppycoder/ngms/accounts/app/server"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// start prometheus metrics handler
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":3002", nil)
	grpclog.Info("prometheus metrics handler starting [::]:3002")

	err := run()
	if err != nil {
		grpclog.Errorf("server was shutdown with errors: %v", err)
		os.Exit(1)
	}
}

func run() error {
	// Application context
	ctx := context.Background()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcAddr("tcp", ":3001"),
		grapiserver.WithGatewayAddr("tcp", ":3000"),
		grapiserver.WithServers(
			server.NewAccountServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
