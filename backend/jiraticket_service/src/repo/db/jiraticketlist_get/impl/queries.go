package impl

const (
	getJiraTickets = `
		select
			j.id,
			j.title,
			j.status,
			j.username
		from
			jira j
	`
)
