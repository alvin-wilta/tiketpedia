package usecase

import (
	"context"

	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/model"
)

type Usecase interface {
	CreateJiraTicket(ctx context.Context, req model.CreateJiraTicketRequest) (model.CreateJiraTicketResponse, error)
	GetOneJiraTicket(ctx context.Context, req model.GetJiraTicketRequest) (model.GetJiraTicketResponse, error)
	GetJiraTicketList(ctx context.Context) (model.GetJiraTicketListResponse, error)
	SetJiraTicketStatus(ctx context.Context, req model.SetJiraTicketStatusRequest) (model.SetJiraTicketStatusResponse, error)
}
