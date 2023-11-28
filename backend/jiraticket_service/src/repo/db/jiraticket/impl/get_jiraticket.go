package impl

import (
	"context"

	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/model"
)

func (r *repository) GetJiraTicket(ctx context.Context, getJiraTicketRequest model.GetJiraTicketRequest) (model.GetJiraTicketResponse, error) {
	// TODO: Implement
	return model.GetJiraTicketResponse{}, nil
}
