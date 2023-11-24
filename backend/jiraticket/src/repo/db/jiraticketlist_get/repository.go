package jiraticketlist_get

import (
	"context"

	"github.com/alvin-wilta/tiketpedia/backend/jiraticket/src/model"
)

type Repository interface {
	GetJiraTicketList(ctx context.Context) (model.GetJiraTicketListResponse, error)
}
