package model

type JiraTicket struct {
	TicketId    int32
	Title       string
	Description string
	UserName    string
	Status      string
}

type JiraTicketShort struct {
	TicketId int32
	Title    string
	Status   string
}

type CreateJiraTicketRequest struct {
	Title       string
	Description string
	UserName    string
}

type CreateJiraTicketResponse struct {
	TicketId int32
	Status   string
}

type GetJiraTicketRequest struct {
	TicketId int32
}

type GetJiraTicketResponse struct {
	JiraTicket JiraTicket
}

type GetJiraTicketListResponse struct {
	JiraTicketList []JiraTicketShort
}

type SetJiraTicketStatusRequest struct {
	TicketId int32
	Status   string
}

type SetJiraTicketStatusResponse struct {
	JiraTicket JiraTicket
	Success    bool
}
