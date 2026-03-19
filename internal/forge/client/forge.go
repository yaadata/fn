package client

import (
	"context"

	. "github.com/yaadata/optionsgo/core"

	forgedomain "github.com/yaadata/fn/internal/forge/domain"
)

type Forge interface {
	GetNotifications(ctx context.Context, cfg ...forgedomain.NotificationGetOptions) Result[forgedomain.ForgeNotifications]
	GetMergeProposals(ctx context.Context, repository forgedomain.Repository) Result[forgedomain.MergeProposals]
	GetIssues(ctx context.Context) Result[forgedomain.Issues]
	OpenIssue(ctx context.Context, issue forgedomain.Issue) bool
	OpenMergeProposal(ctx context.Context, mp forgedomain.MergeProposal) bool
}
