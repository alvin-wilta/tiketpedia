package main

import (
	"log"
	"net"

	"github.com/alvin-wilta/tiketpedia/backend/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcServer := &GRPCserver{}
	proto.RegisterJiraServiceServer(s, grpcServer)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
