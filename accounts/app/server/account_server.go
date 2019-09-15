package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "accounts/api"
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

func (s *accountServiceServerImpl) ListAccounts(ctx context.Context, req *api_pb.ListAccountsRequest) (*api_pb.ListAccountsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *accountServiceServerImpl) GetAccount(ctx context.Context, req *api_pb.GetAccountRequest) (*api_pb.Account, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *accountServiceServerImpl) CreateAccount(ctx context.Context, req *api_pb.CreateAccountRequest) (*api_pb.Account, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *accountServiceServerImpl) UpdateAccount(ctx context.Context, req *api_pb.UpdateAccountRequest) (*api_pb.Account, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *accountServiceServerImpl) DeleteAccount(ctx context.Context, req *api_pb.DeleteAccountRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
