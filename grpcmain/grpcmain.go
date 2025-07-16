package grpcmain

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"mbooke/grpc-server/global"
	pb "mbooke/grpc-server/proto/location"
	pbm "mbooke/grpc-server/proto/message"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pbm.UnimplementedStreamServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SetLocation(_ context.Context, in *pb.LocationRequest) (*pb.LocationReply, error) {
	log.Printf("Received: %s - %f - %f", in.Id, in.Longitude, in.Latitude)
	global.StreamServerMap.Lock()
	srv := global.StreamServerMap.M[in.GetId()]
	resp := pbm.Response{Result: "Test Send Pesan"}

	if srv != nil {
		log.Println("Send data grpc stream")
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}

	global.StreamServerMap.Unlock()
	return &pb.LocationReply{Message: "LocationId " + in.GetId(), Status: 1}, nil
}

func (s *server) SendMessage(_ context.Context, in *pbm.MessageRequest) (*pbm.MessageResponse, error) {
	global.StreamServerMap.Lock()
	srv := global.StreamServerMap.M[in.GetId()]
	resp := pbm.Response{Result: in.Message}

	if srv != nil {
		log.Printf("Send data grpc stream to %s", in.GetId())
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	} else {
		log.Printf("Client %s not found", in.GetId())
	}

	global.StreamServerMap.Unlock()
	return &pbm.MessageResponse{Result: "Success"}, nil
}

func Grpcmain() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pbm.RegisterStreamServiceServer(s, &server{})
	log.Printf("server stream listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
