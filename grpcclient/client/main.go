// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pbm "mbooke/grpc-server/proto/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "client1"
)

var (
	addr = flag.String("addr", "103.56.149.92:50051", "the address to connect to")
	id   = flag.String("name", defaultName, "Name to greet")
	// longitude = 1213123.3
	// latitude  = 121212.3
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbm.NewStreamServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &pbm.MessageRequest{Id: *id, Message: "test Send Data Client 1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status %s", r.GetResult())
}
