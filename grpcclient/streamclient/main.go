package main

import (
	"context"
	"io"
	"log"
	"math/rand"

	pb "mbooke/grpc-server/proto/message"

	"time"

	"google.golang.org/grpc"
)

func main() {

	rand.Seed(time.Now().Unix())

	// dial server
	conn, err := grpc.Dial("103.56.149.92:50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{Id: "client1"}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed by server")
			break
		}
		if err != nil {
			log.Fatalf("can not receive %v", err)
		}
		log.Printf("Resp received: %s", resp.Result)
	}
}
