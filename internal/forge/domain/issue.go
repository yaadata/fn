package forgedomain

type Issue struct {
	Author     string
	Id         uint
	Context    IssueContext
	Repository Repository
	Status     Value[IssueStatus]
	Title      string
}
