package server

import (
	"context"
	"google.golang.org/grpc/grpclog"

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
	id := req.AccountId
	if "" == id || "0" == id {
		grpclog.Infof("invalid account id: %s", id)
		return nil, status.Error(codes.InvalidArgument, "Invalid account id")
	}

	acc, err := repo.GetAccountById(ctx, id)
	if err != nil {
		grpclog.Infof("unable to retrieve account summary for %s", id)
		return nil, status.Error(codes.Unavailable, "Unable to retrieve account")
	}

	return acc, nil
}
