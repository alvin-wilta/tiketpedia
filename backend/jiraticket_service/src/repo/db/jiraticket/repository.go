package jiraticket

import (
	"context"

	std "github.com/alvin-wilta/tiketpedia/backend/common/model"
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/model"
)

type Repository interface {
	CreateJiraTicket(ctx context.Context, createJiraTicketRequest model.CreateJiraTicketRequest) (std.StandardResponse, error)
	GetJiraTicket(ctx context.Context, getJiraTicketRequest model.GetJiraTicketRequest) (model.GetJiraTicketResponse, error)
	SetJiraTicketStatus(ctx context.Context, setJiraTicketStatusRequest model.SetJiraTicketStatusRequest) (std.StandardResponse, error)
	GetJiraTicketList(ctx context.Context) (model.GetJiraTicketListResponse, error)
}
