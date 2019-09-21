package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/myzhan/boomer"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

var client api_pb.AccountServiceClient

const IdFile = "ids.txt"

func idsFromFile(fname string) ([]string, int) {
	log.Print("reading ids from file...")

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Print("cannot open ids.txt file")
		return nil, 0
	}

	lines := strings.Split(string(content), "\n")
	return lines, len(lines)
}

func readIdsFromDB() ([]string, int) {
	log.Print("reading ids from database...")

	ctx := context.Background()

	dbURI := os.Getenv("DBURI")
	if dbURI == "" {
		dbURI = "mongodb://dev:dev@127.0.0.1:27017/dev?authSource=dev&authMechanism=SCRAM-SHA-256"
	}
	r, _ := regexp.Compile("mongodb://.*/(.+)\\?")
	m := r.FindStringSubmatch(dbURI)
	if len(m) <= 1 {
		log.Fatalf("Cannot find dbname from DBURI %s", dbURI)
	}
	dbname := m[1]

	clientOpts := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbname)
	defer client.Disconnect(ctx)

	projection := bson.D{
		{"accountId", 1},
		{"_id", 0},
	}
	cur, err := db.Collection("accounts").Find(
		ctx, bson.D{},
		options.Find().SetProjection(projection),
	)
	if err != nil {
		log.Fatal(err)
	}

	var ids []string
	for cur.Next(ctx) {
		var doc struct {
			AccountId string
		}
		err := cur.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, doc.AccountId)
	}

	return ids, len(ids)
}

func randId() func() string {
	// try to read from ids.txt file first
	// if nothing is read then read from database
	ids, size := idsFromFile(IdFile)
	if size == 0 {
		ids, size = readIdsFromDB()
	}
	log.Printf("using a pool of %d ids", size)

	rs := rand.NewSource(time.Now().UnixNano())
	rr := rand.New(rs)

	return func() string {
		id := strings.TrimSpace(ids[rr.Intn(size)])
		if id == "" {
			log.Print("Got an empty id, perhaps there's some bugs here")
			id = "0000"
		}
		return id
	}
}

func setupGrpcApi(name, addr string, randId func() string) func() {
	log.Print("setting up gRPC API test")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()

	client = api_pb.NewAccountServiceClient(conn)

	return func() {
		ctx := context.Background()

		start := time.Now()
		id := randId()
		r, err := client.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: id})
		elapsed := time.Since(start)

		if err != nil || r == nil {
			log.Printf("account %s got error %s", id, err)
			boomer.RecordFailure(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
		} else {
			boomer.RecordSuccess(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				int64(10),
			)
		}
	}
}

func setupRestApi(name, baseUrl string, randId func() string) func() {
	log.Print("setting up REST API test")

	return func() {
		id := randId()
		url := baseUrl + "/" + id
		// log.Print(url)

		start := time.Now()
		r, err := http.Get(url)
		elapsed := time.Since(start)

		if err != nil || r == nil {
			log.Printf("HTTP GET to %s returned %s", url, err.Error())
			boomer.RecordFailure(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
			return
		}

		defer r.Body.Close()
		_, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Reading response for HTTP GET %s, got error %s", url, err.Error())
			boomer.RecordFailure(
				"http",
				name,
				elapsed.Nanoseconds()/int64(time.Millisecond),
				err.Error(),
			)
			return
		}

		boomer.RecordSuccess(
			"http",
			name,
			elapsed.Nanoseconds()/int64(time.Millisecond),
			int64(10),
		)
	}
}

func testGrpc(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api_pb.NewAccountServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: "100-1122-5577-891"})
	if err != nil {
		log.Printf("error calling API: %v", err)
	}
	log.Printf("Greeting: %s", r)
}

func main() {
	var proto string
	flag.StringVar(&proto, "proto", "grpc", "protocol to test: rest or grpc")
	testOnly := flag.Bool("test", false, "call gRPC API then exit, will not run as slave")

	if !flag.Parsed() {
		flag.Parse()
	}

	restAddr := os.Getenv("RESTSVC_ADDR")
	if restAddr == "" {
		restAddr = "http://[::]:3000/accounts"
	}

	grpcAddr := os.Getenv("GRPCSVC_ADDR")
	if grpcAddr == "" {
		grpcAddr = "[::]:3001"
	}

	testGrpc(grpcAddr)

	if *testOnly {
		log.Print("test only mode, exiting...")
		os.Exit(0)
	}

	var task *boomer.Task
	if proto == "grpc" {
		task = &boomer.Task{
			Name:   "gprc",
			Weight: 100, // The weight is used to distribute goroutines over multiple tasks.
			Fn:     setupGrpcApi("gRPC account api", grpcAddr, randId()),
		}
	} else {
		task = &boomer.Task{
			Name:   "rest",
			Weight: 100,
			Fn:     setupRestApi("REST account api", restAddr, randId()),
		}
	}

	boomer.Run(task)
}
