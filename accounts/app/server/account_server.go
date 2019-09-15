package server

import (
	"context"

	// "github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
	repo "github.com/sloppycoder/ngms/accounts/app/repo"
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
	if req.AccountId == "" || req.AccountId == "0" {
		return nil, status.Error(codes.InvalidArgument, "Invalid account id")
	}

	acc, err := repo.GetAccountById(ctx, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "Unable to retrieve account")
	}

	return acc, nil
}
