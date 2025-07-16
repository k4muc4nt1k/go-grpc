package grpcmain

import (
	"flag"
	"fmt"
	"log"
	"mbooke/grpc-server/global"
	pb "mbooke/grpc-server/proto/message"
	"net"

	"google.golang.org/grpc"
)

var (
	portStream = flag.Int("portStream", 50005, "The server port")
)

type serverstream struct {
	pb.UnimplementedStreamServiceServer
}

// func myUnaryInterceptor(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (interface{}, error) {
// 	// Your logic here
// 	return handler(ctx, req)
// }

func (s serverstream) FetchResponse(in *pb.Request, srv pb.StreamService_FetchResponseServer) error {

	log.Printf("fetch response for id : %s", in.Id)

	// Set srv in global StreamServerMap using in.Id as key
	key := in.Id
	global.StreamServerMap.Lock()
	global.StreamServerMap.M[key] = srv
	global.StreamServerMap.Unlock()

	resp := pb.Response{Result: "Registration Success"}
	if err := srv.Send(&resp); err != nil {
		log.Printf("send error %v", err)
		return err
	}

	<-srv.Context().Done()
	log.Printf("stream for id %s closed by client", in.Id)
	global.StreamServerMap.Lock()
	delete(global.StreamServerMap.M, key)
	global.StreamServerMap.Unlock()
	return nil
}

func Streammain() {

	// tlsCredentials, err := auth.LoadTLSCredentials()
	// if err != nil {
	// 	log.Fatal("cannot load TLS credentials: ", err)
	// }

	// create listiner
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portStream))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portStream))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// s := grpc.NewServer(
	// 	grpc.Creds(tlsCredentials),
	// 	grpc.UnaryInterceptor(myUnaryInterceptor),
	// )

	// pb.RegisterStreamServiceServer(s, serverstream{})

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, serverstream{})

	log.Printf("server listening at %v", lis.Addr())
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
