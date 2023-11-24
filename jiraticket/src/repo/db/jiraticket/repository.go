package jiraticket

import (
	"context"

	"github.com/alvin-wilta/tiketpedia/jiraticket/src/model"
)

type Repository interface {
	CreateJiraTicket(ctx context.Context, createJiraTicketRequest model.CreateJiraTicketRequest) (model.StandardResponse, error)
	GetJiraTicket(ctx context.Context, getJiraTicketRequest model.GetJiraTicketRequest) (model.GetJiraTicketResponse, error)
	SetJiraTicketStatus(ctx context.Context, setJiraTicketStatusRequest model.SetJiraTicketStatusRequest) (model.StandardResponse, error)
}
