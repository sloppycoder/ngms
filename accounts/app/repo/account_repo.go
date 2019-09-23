package repo

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/grpclog"
	"os"
	"time"
)

type Balance struct {
	Amount      float64
	Type        string
	CreditFlag  bool
	LastUpdated primitive.DateTime `bson:"last_updated, omitempty"`
}

type Account struct {
	Id                primitive.ObjectID `bson:"_id, omitempty"`
	AccountId         string
	Nickname          string
	ProdCode          string
	ProdName          string
	Currency          string
	Servicer          string
	Status            string
	StatusLastUpdated primitive.DateTime `bson:"status_last_updated, omitempty"`
	Balances          []*Balance
}

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
		Name: "get_account_success_histogram",
		Help: "histogram for successful fetch account from database",
		// used to store microseconds
		Buckets: []float64{100, 250, 500, 1_000, 2_500, 5_000, 10_000, 25_000, 50_000, 100_000, 2_500_000},
	})
)

func GetAccountById(ctx context.Context, id string) (*api_pb.Account, error) {

	db := db(ctx)
	start := time.Now()
	cur := db.Collection("accounts").FindOne(ctx, bson.M{"accountId": id})

	if err := cur.Err(); err != nil {
		getAccountFailure.Inc()
		return nil, err
	}

	var acc Account
	err := cur.Decode(&acc)
	if err != nil {
		getAccountFailure.Inc()
		return nil, err
	}

	account, err := mapAccount(&acc)
	if err != nil {
		grpclog.Info("Mapping account got error %v", err)
	}

	getAccountSuccess.Inc()
	elapsed := time.Since(start)
	getAccountHisto.Observe(float64(elapsed / time.Microsecond))

	return account, nil
}

func mapAccount(acc *Account) (*api_pb.Account, error) {
	var account api_pb.Account

	copier.Copy(&account, &acc)
	account.Id = acc.Id.String()[10:34] // strip off the ObjectId("...")

	var err error
	account.StatusLastUpdated, err = ptypes.TimestampProto(acc.StatusLastUpdated.Time())

	for i, bal := range acc.Balances {
		pbtimestamp, err2 := ptypes.TimestampProto(bal.LastUpdated.Time())
		account.Balances[i].LastUpdated = pbtimestamp
		// no sure what we can do about incorrect timestamp here,
		// just pass back some indicator
		if err2 != nil && err == nil {
			err = err2
		}
	}

	return &account, err
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
