package forgedomain

import . "github.com/yaadata/optionsgo"

type Issue struct {
	Author     string
	Id         IssueId
	Context    IssueContext
	Repository Option[Repository]
	Status     Value[IssueStatus]
	Title      string
}
