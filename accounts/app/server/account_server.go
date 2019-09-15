package server

import (
	"context"

	// "github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

// AccountServiceServer is a composite interface of api_pb.AccountServiceServer and grapiserver.Server.
type AccountServiceServer interface {
	api_pb.AccountServiceServer
	grapiserver.Server
}

// NewAccountServiceServer creates a new AccountServiceServer instance.
func NewAccountServiceServer() AccountServiceServer {
	return &accountServiceServerImpl{}
}

type accountServiceServerImpl struct {
}

func (s *accountServiceServerImpl) GetAccount(ctx context.Context, req *api_pb.GetAccountRequest) (*api_pb.Account, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
