package forgedomain

type IssueStatus string

const (
	IssueStatusNone   IssueStatus = ""
	IssueStatusOpen   IssueStatus = "open"
	IssueStatusClosed IssueStatus = "closed"
)
