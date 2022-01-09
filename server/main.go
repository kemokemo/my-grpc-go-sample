package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/kemokemo/my-grpc-sample"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type echoService struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}

func init() {
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Println("failed to listen: ", err)
		return 1
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &echoService{})
	log.Println("server listening at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: ", err)
		return 1
	}

	return 0
}
