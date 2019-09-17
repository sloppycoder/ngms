package server

import (
	"context"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/grpclog"
	"strconv"
)

const (
	DBName = "ngmsdev"
	DBURI  = "mongodb://dev:pass@127.0.0.1:27017/ngmsdev?authSource=ngmsdev&authMechanism=SCRAM-SHA-256"
)

func GetAccountById(ctx context.Context, id string) (*api_pb.Account, error) {
	db := db(ctx)
	cur := db.Collection("accounts").FindOne(ctx, bson.M{"accountId": id})
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// TODO:
	//  rewrite decodeRaw logic with a more concise method
	// this code is ugly!

	raw, err := cur.DecodeBytes()
	var ob map[string]interface{}
	_ = bson.Unmarshal(raw, &ob)

	balances := make([]*api_pb.Balance, 2)

	if ob["balances"] != nil {
		bt, err := strconv.Atoi(ob["balances"].(primitive.A)[0].(map[string]interface{})["type"].(string))
		if err == nil {

			balances[0] = &api_pb.Balance{
				Amount:     ob["balances"].(primitive.A)[0].(map[string]interface{})["amount"].(float64),
				Type:       api_pb.Balance_Type(bt),
				CreditFlag: ob["balances"].(primitive.A)[0].(map[string]interface{})["creditFlag"].(bool),
			}
		}

		bt2, err := strconv.Atoi(ob["balances"].(primitive.A)[1].(map[string]interface{})["type"].(string))
		if err == nil {
			balances[1] = &api_pb.Balance{
				Amount:     ob["balances"].(primitive.A)[1].(map[string]interface{})["amount"].(float64),
				Type:       api_pb.Balance_Type(bt2),
				CreditFlag: ob["balances"].(primitive.A)[1].(map[string]interface{})["creditFlag"].(bool),
			}
		}
	}

	account := api_pb.Account{
		AccountId: ob["accountId"].(string),
		ProdCode:  ob["prodCode"].(string),
		ProdName:  ob["prodName"].(string),
		Balances:  balances,
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func db(ctx context.Context) *mongo.Database {
	clientOpts := options.Client().ApplyURI(DBURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		grpclog.Error(err)
	}

	return client.Database(DBName)
}
