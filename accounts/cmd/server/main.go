package main

import (
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/sloppycoder/ngms/accounts/app/server"
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

func run() error {
	// Application context
	ctx := context.Background()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcAddr("tcp", "0.0.0.0:3001"),
		grapiserver.WithGatewayAddr("tcp", "0.0.0.0:3000"),
		grapiserver.WithServers(
			server.NewAccountServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
