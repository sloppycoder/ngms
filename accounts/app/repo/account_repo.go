package server

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

const (
	DBName = "ngmsdev"
	DBURI  = "mongodb://dev:pass@127.0.0.1:27017/ngmsdev?authSource=ngmsdev&authMechanism=SCRAM-SHA-256"
)

func GetAccountById(ctx context.Context, id string) (*api_pb.Account, error) {

	db := db(ctx)
	res := db.Collection("accounts").FindOne(ctx, bson.M{"accountId": id})
	if err := res.Err(); err != nil {
		return nil, err
	}
	//acc := &api_pb.Account{
	//	Id:        "100",
	//	AccountId: "100-XXX-XXXX-1234",
	//	ProdCode:  "2001",
	//	ProdName:  "Super Saver Bonus Account",
	//}
	var account api_pb.Account
	err := res.Decode(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func db(ctx context.Context) *mongo.Database {
	clientOpts := options.Client().ApplyURI(DBURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(DBName)
}
