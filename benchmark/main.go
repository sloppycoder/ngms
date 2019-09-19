package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/myzhan/boomer"
	api_pb "github.com/sloppycoder/ngms/accounts/api"
)

var client api_pb.AccountServiceClient
var ids *[]string
var totalIds int
var rr *rand.Rand

func randId() string {
	i := rr.Intn(totalIds)
	if i >=0 && i <= totalIds -1 {
		return (*ids)[i]
	}
	return (*ids)[0]
}

func invokeGrpcApi(){
	ctx := context.Background()

	start := time.Now()
	r, err := client.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: randId()})
	elapsed := time.Since(start)

	if err != nil || r == nil {
		boomer.RecordFailure("http", "invokeGrpcApi", elapsed.Nanoseconds()/int64(time.Millisecond), "cannot retrieve account detail")
	}

	/*
	   Report your test result as a success, if you write it in locust, it will looks like this
	   events.request_success.fire(request_type="http", name="foo", response_time=100, response_length=10)
	*/
	boomer.RecordSuccess("http", "invokeGrpcApi", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}

func invokeRestApi(){
	start := time.Now()
	r, err := http.Get("http://127.0.0.1:3000/accounts/" + randId())
	elapsed := time.Since(start)

	if err != nil || r == nil {
		boomer.RecordFailure("http", "invokeRestApi", elapsed.Nanoseconds()/int64(time.Millisecond), "cannot retrieve account detail")
	}

	defer r.Body.Close()
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		boomer.RecordFailure("http", "invokeRestApi", elapsed.Nanoseconds()/int64(time.Millisecond), "unable to read response")
	}
	/*
	   Report your test result as a success, if you write it in locust, it will looks like this
	   events.request_success.fire(request_type="http", name="foo", response_time=100, response_length=10)
	*/
	boomer.RecordSuccess("http", "invokeRestApi", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
}

func main(){
	var proto string
	flag.StringVar(&proto, "proto", "grpc", "protocol to test: rest or grpc")

	if !flag.Parsed() {
		flag.Parse()
	}

	content, err := ioutil.ReadFile("ids.txt")
	if err != nil {
		log.Fatal("cannot open ids.txt file.")
	}

	lines := strings.Split(string(content), "\n")
	ids = &lines
	totalIds = len(lines)
	log.Printf("using a pool of %d ids", totalIds)

	rr = rand.New(rand.NewSource(time.Now().UnixNano()))

	restTask := &boomer.Task{
		Name: "invokeRestApi",
		// The weight is used to distribute goroutines over multiple tasks.
		Weight: 100,
		Fn: invokeRestApi,
	}

	grpcTask := &boomer.Task{
		Name: "invokeGrpcApi",
		// The weight is used to distribute goroutines over multiple tasks.
		Weight: 100,
		Fn: invokeGrpcApi,
	}

	if proto == "grpc" {
		conn, err := grpc.Dial("[::]:3001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client = api_pb.NewAccountServiceClient(conn)
		boomer.Run(grpcTask)
	} else {
		boomer.Run(restTask)
	}
}