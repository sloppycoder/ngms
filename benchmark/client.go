package main

import (
	"context"
	"log"
	"os"
	"time"

	api_pb "github.com/sloppycoder/ngms/accounts/api"
	"google.golang.org/grpc"
)

func main() {
	addr := os.Getenv("GRPCSVC_ADDR")
	if addr == "" {
		addr = "[::]:3001"
	}

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
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
