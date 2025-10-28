package main

import (
	"context"
	"log"
	"net"

	pb "github.com/IdentityStolen/RiskOS/grpc"
	"google.golang.org/grpc"
)

type checkerServer struct {
	pb.UnimplementedRiskCheckerServer
}

func (s *checkerServer) TransScrutiny(ctx context.Context, in *pb.TransRequest) (*pb.TransResponce, error) {
	// return nil, fmt.Errorf("failed to process data:")
	log.Printf("context: %q", ctx)
	return &pb.TransResponce{IsOk: true}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8085")

	if err != nil {
		log.Fatalf("Error creating listener: %q", err)
	}

	serverRegister := grpc.NewServer()
	service := &checkerServer{}
	pb.RegisterRiskCheckerServer(serverRegister, service)
	log.Printf("starting server and serving requests...")

	err = serverRegister.Serve(listener)
	if err != nil {
		log.Fatalf("Something went wrong: %q", err)
	}
}
