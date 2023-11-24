package impl

const (
	getJiraTicketById = `
		select
			j.id,
			j.title,
			j.description,
			j.status,
			j.username
		from
			jira j
		where 
			j.id = $1
		limit 1;
	`
)
