package repo

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/grpclog"
	"os"
	"strconv"
	"time"
)

var _db *mongo.Database
var (
	getAccountSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_account_success_total",
		Help: "total accounts successfully fetched from database",
	})

	getAccountFailure = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_account_failure_total",
		Help: "total failed attempts for fetching account from database",
	})

	getAccountHisto = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "get_account_success_histogram",
		Help:    "histogram for successful fetch account from database",
		Buckets: []float64{5, 20, 50, 100, 200, 500, 1000, 2000, 5000},
	})
)

func GetAccountById(ctx context.Context, id string) (*api_pb.Account, error) {
	db := db(ctx)
	start := time.Now()
	cur := db.Collection("accounts").FindOne(ctx, bson.M{"accountId": id})
	elapsed := time.Since(start)
	getAccountHisto.Observe(float64(elapsed) / 1_000) // millisecond

	if err := cur.Err(); err != nil {
		getAccountFailure.Inc()
		return nil, err
	}

	account, err := fetchAccount(cur)
	if err != nil {
		getAccountFailure.Inc()
		return nil, err
	}

	getAccountSuccess.Inc()
	return account, nil
}

func db(ctx context.Context) *mongo.Database {
	if _db != nil {
		return _db
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		dbName = "dev"
	}

	dbURI := os.Getenv("DBURI")
	if dbURI == "" {
		dbURI = "mongodb://dev:dev@127.0.0.1:27017/dev?authSource=dev&authMechanism=SCRAM-SHA-256"
	}
	grpclog.Infof("DBURI=%s", dbURI)

	clientOpts := options.Client().ApplyURI(dbURI).SetMinPoolSize(10).SetMaxPoolSize(100)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		grpclog.Error(err)
	}

	_db = client.Database(dbName)
	// TODO: when do I disconnect?
	return _db
}

func xfetchAccount(cur *mongo.SingleResult) (*api_pb.Account, error) {
	var acc api_pb.Account
	err := cur.Decode(&acc)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func fetchAccount(cur *mongo.SingleResult) (*api_pb.Account, error) {
	// TODO:
	//  rewrite decodeRaw logic with a more concise method
	// this code is ugly!

	raw, err := cur.DecodeBytes()
	if err != nil {
		return nil, err
	}

	var ob map[string]interface{}
	_ = bson.Unmarshal(raw, &ob)

	acc := api_pb.Account{
		AccountId: ob["accountId"].(string),
		ProdCode:  ob["prodCode"].(string),
		ProdName:  ob["prodName"].(string),
		Balances:  nil,
	}

	if ob["balances"] != nil {
		balances := make([]*api_pb.Balance, 2)

		bt1, err1 := strconv.Atoi(ob["balances"].(primitive.A)[0].(map[string]interface{})["type"].(string))
		if err1 == nil {
			balances[0] = &api_pb.Balance{
				Amount:     ob["balances"].(primitive.A)[0].(map[string]interface{})["amount"].(float64),
				Type:       api_pb.Balance_Type(bt1),
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

		acc.Balances = balances
	}

	return &acc, nil
}
