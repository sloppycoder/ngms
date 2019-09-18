package main

import (
	"context"
	"log"
	"time"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
	"google.golang.org/grpc"
)

const (
	address     = "[::]:3001"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api_pb.NewAccountServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.GetAccount(ctx, &api_pb.GetAccountRequest{AccountId: "100-1234-5577-891"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
