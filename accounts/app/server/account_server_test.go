package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

func Test_AccountServiceServer_GetAccount(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.GetAccountRequest{AccountId: "100"}

	resp, err := svr.GetAccount(ctx, req)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func Test_AccountServiceServer_GetAccount_Invalid_AccountId(t *testing.T) {
	svr := NewAccountServiceServer()

	ctx := context.Background()
	req := &api_pb.GetAccountRequest{AccountId: ""}

	_, err := svr.GetAccount(ctx, req)

	assert.NotNil(t, err)
}
