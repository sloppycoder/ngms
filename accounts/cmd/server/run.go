package main

import (
	"context"
	"github.com/sloppycoder/ngms/accounts/app/server"

	"github.com/izumin5210/grapi/pkg/grapiserver"
)

func run() error {
	// Application context
	ctx := context.Background()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewAccountServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
