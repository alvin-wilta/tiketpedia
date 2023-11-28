package main

import (
	"context"
	"log"

	pb "github.com/alvin-wilta/tiketpedia/backend/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NOTE: gRPC and NSQ are implemented here

type GRPCserver struct {
	pb.UnimplementedJiraServiceServer
}

func (s *GRPCserver) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	var res pb.HealthCheckResponse
	res.Ok = true
	log.Println("Health check succeeded")
	// TODO: Implement
	return &res, nil
}

func (s *GRPCserver) CreateJiraTicket(ctx context.Context, createJiraTicketRequest *pb.CreateJiraTicketRequest) (*pb.CreateJiraTicketResponse, error) {
	var res pb.CreateJiraTicketResponse
	// TODO: Implement
	return &res, nil
}

func (s *GRPCserver) GetOneJiraTicket(ctx context.Context, getOneJiraTicketRequest *pb.GetOneJiraTicketRequest) (*pb.Ticket, error) {
	var res pb.Ticket
	// TODO: Implement
	return &res, nil
}

func (s *GRPCserver) GetAllJiraTicket(ctx context.Context, empty *emptypb.Empty) (*pb.GetJiraTicketListResponse, error) {
	var res pb.GetJiraTicketListResponse
	// TODO: Implement
	return &res, nil
}
