package server

import (
	"context"
	"testing"

	api_pb "accounts/api"
)

func Test_AccountServiceServer_ListAccounts(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.ListAccountsRequest{}

	resp, err := svr.ListAccounts(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_AccountServiceServer_GetAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.GetAccountRequest{}

	resp, err := svr.GetAccount(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_AccountServiceServer_CreateAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateAccountRequest{}

	resp, err := svr.CreateAccount(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_AccountServiceServer_UpdateAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateAccountRequest{}

	resp, err := svr.UpdateAccount(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_AccountServiceServer_DeleteAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteAccountRequest{}

	resp, err := svr.DeleteAccount(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
