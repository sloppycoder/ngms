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

	resp, err := svr.GetAccount(ctx,
		&api_pb.GetAccountRequest{
			AccountId: "100-1122-5577-891",
		})

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, resp.ProdCode, "1220")
}

func Test_AccountServiceServer_GetAccount_Invalid_AccountId(t *testing.T) {
	svr := NewAccountServiceServer()
	ctx := context.Background()

	_, err := svr.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: ""})

	assert.NotNil(t, err)
}
