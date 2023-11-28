package impl

import (
	"context"

	std "github.com/alvin-wilta/tiketpedia/backend/common/model"
	"github.com/alvin-wilta/tiketpedia/backend/jiraticket_service/src/model"
)

func (r *repository) SetJiraTicketStatus(ctx context.Context, setJiraTicketStatusRequest model.SetJiraTicketStatusRequest) (std.StandardResponse, error) {
	// TODO: Implement
	return std.StandardResponse{}, nil
}
