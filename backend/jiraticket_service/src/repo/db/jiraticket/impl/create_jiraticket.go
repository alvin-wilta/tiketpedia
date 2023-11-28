package impl

import (
	"context"

	std "github.com/alvin-wilta/tiketpedia/backend/common/model"
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/model"
)

func (repo *repository) CreateJiraTicket(ctx context.Context, createJiraTicketRequest model.CreateJiraTicketRequest) (std.StandardResponse, error) {
	// TODO: Implement
	return std.StandardResponse{}, nil
}
