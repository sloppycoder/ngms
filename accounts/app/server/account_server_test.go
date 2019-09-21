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
			AccountId: "111",
		})

	assert.NotNil(t, resp)
	assert.NotNil(t, resp.StatusLastUpdated)
	assert.EqualValues(t, len(resp.Balances), 2)
	assert.Nil(t, err)
	assert.EqualValues(t, resp.ProdCode, "1002")
}

func Test_AccountServiceServer_GetAccount_Invalid_AccountId(t *testing.T) {
	svr := NewAccountServiceServer()
	ctx := context.Background()

	_, err := svr.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: ""})

	assert.NotNil(t, err)
}
