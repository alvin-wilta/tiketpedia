package jiraticket

import "context"

type Repository interface {
	GetJiraTicketExist(ctx context.Context, key string) (bool, error)
	GetJiraTicket(ctx context.Context, key string) (bool, error)
	SetJiraTicket(ctx context.Context, key string, value interface{}) error
}
