package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/myzhan/boomer"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

var client api_pb.AccountServiceClient

func randId() func() string {
	content, err := ioutil.ReadFile("ids.txt")
	if err != nil {
		log.Fatal("cannot open ids.txt file.")
	}

	lines := strings.Split(string(content), "\n")
	size := len(lines)
	log.Printf("using a pool of %d ids", size)

	rs := rand.NewSource(time.Now().UnixNano())
	rr := rand.New(rs)

	return func() string {
		id := strings.TrimSpace(lines[rr.Intn(size)])
		if id == "" {
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
			boomer.RecordFailure("http", name, elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())

		} else {
			boomer.RecordSuccess("http", name, elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
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
			boomer.RecordFailure("http", name, elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
			return
		}

		defer r.Body.Close()
		_, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Reading response for HTTP GET %s, got error %s", url, err.Error())
			boomer.RecordFailure("http", name, elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
			return
		}

		boomer.RecordSuccess("http", name, elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
	}
}

func main() {
	var proto string
	flag.StringVar(&proto, "proto", "grpc", "protocol to test: rest or grpc")

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
