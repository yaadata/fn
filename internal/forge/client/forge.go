package client

import (
	"context"

	. "github.com/yaadata/optionsgo/core"

	forgedomain "github.com/yaadata/fn/internal/forge/domain"
)

type Forge interface {
	GetNotifications(ctx context.Context, cfg ...forgedomain.NotificationGetOptions) Result[forgedomain.ForgeNotifications]
	GetMergeProposals(ctx context.Context) Result[forgedomain.MergeProposals]
	GetIssues(ctx context.Context) Result[forgedomain.Issues]
}
