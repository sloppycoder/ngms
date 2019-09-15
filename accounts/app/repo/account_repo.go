package server

import (
	"context"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

func GetAccountById(ctx context.Context, id string) (*api_pb.Account, error) {
	acc := &api_pb.Account{
		Id:        "100",
		AccountId: "100-XXX-XXXX-1234",
		ProdCode:  "2001",
		ProdName:  "Super Saver Bonus Account",
	}
	return acc, nil
}
