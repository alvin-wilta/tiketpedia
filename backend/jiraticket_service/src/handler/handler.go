package handler

import (
	"context"

	pb "github.com/alvin-wilta/tiketpedia/backend/proto"
)

type Handler interface {
	CreateJiraTicket(ctx context.Context, req *pb.CreateJiraTicketRequest) (*pb.CreateJiraTicketResponse, error)
	GetOneJiraTIcket(ctx context.Context, req *pb.GetOneJiraTicketRequest) (*pb.CreateJiraTicketResponse, error)
	GetJiraTicketList(ctx context.Context) (*pb.GetJiraTicketListResponse, error)
	SetJiraTicketStatus(ctx context.Context, req *pb.SetJiraTicketStatusRequest) (*pb.SetJiraTicketStatusResponse, error)
}
