package server

import (
	"context"
	"testing"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

func Test_AccountServiceServer_GetAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.GetAccountRequest{
		AccountId: "100",
	}

	resp, err := svr.GetAccount(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_AccountServiceServer_GetAccount_Invalid_AccountId(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.GetAccountRequest{
		AccountId: "",
	}

	_, err := svr.GetAccount(ctx, req)

	if err == nil {
		t.Errorf("should have returend an error %v", err)
	}
}
