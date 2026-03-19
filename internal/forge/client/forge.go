package client

import (
	"context"

	forgedomain "github.com/yaadata/fn/internal/forge/domain"
	. "github.com/yaadata/optionsgo"
)

type Forge interface {
	GetNotifications(ctx context.Context, cfg ...forgedomain.NotificationGetOptions) Result[forgedomain.ForgeNotifications]
	GetMergeProposals(ctx context.Context, repository forgedomain.Repository) Result[forgedomain.MergeProposals]
	GetIssues(ctx context.Context) Result[forgedomain.Issues]
	OpenIssue(ctx context.Context, issue forgedomain.Issue) bool
	OpenMergeProposal(ctx context.Context, mp forgedomain.MergeProposal) bool
}
