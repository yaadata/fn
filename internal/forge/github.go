package forge

import (
	"context"
	"strings"

	"github.com/google/go-github/v79/github"
	. "github.com/yaadata/optionsgo"
	"github.com/yaadata/optionsgo/extension"

	"github.com/yaadata/fn/internal/forge/client"
	forgedomain "github.com/yaadata/fn/internal/forge/domain"
)

type githubForge struct {
	client *github.Client
}

var _ client.Forge = (*githubForge)(nil)

func GithubForge(client *github.Client) *githubForge {
	return &githubForge{
		client: client,
	}
}

func (g *githubForge) GetNotifications(ctx context.Context, cfg ...forgedomain.NotificationGetOptions) Result[forgedomain.ForgeNotifications] {
	notifications, _, err := g.client.Activity.ListNotifications(ctx, &github.NotificationListOptions{})
	if err != nil {
		return Err[forgedomain.ForgeNotifications](err)
	}
	var resp forgedomain.ForgeNotifications
	for _, n := range notifications {
		resp = append(resp, &forgedomain.Notification{
			Context: forgedomain.NotificationContext{
				Url:         extension.OptionFromPointer(n.GetSubject().URL),
				Description: extension.OptionFromPointer(n.GetSubject().Title),
			},
			Id:         n.GetID(),
			Kind:       githubNotificationSubjectTypeToKind(n.Subject.Type),
			Repository: githubRepository(n.Repository),
			Trigger:    githubNotificationTrigger(n.GetReason()),
			Status:     githubNotificationStatus(!n.GetUnread()),
			Viewed:     !n.GetUnread(),
		})
	}

	return Ok(resp)
}

func (g *githubForge) GetMergeProposals(ctx context.Context, repository forgedomain.Repository) Result[forgedomain.MergeProposals] {
	pullRequests, _, err := g.client.PullRequests.List(ctx, repository.Group, repository.Project, nil)
	if err != nil {
		return Err[forgedomain.MergeProposals](err)
	}
	var resp forgedomain.MergeProposals
	for _, p := range pullRequests {
		labels := make([]string, 0, len(p.Labels))
		for i, l := range p.Labels {
			labels[i] = l.GetName()
		}
		resp = append(resp, forgedomain.MergeProposal{
			Author: p.GetUser().GetName(),
			Context: forgedomain.MergeProposalContext{
				Body:   p.GetBody(),
				Labels: labels,
				Url:    extension.OptionFromPointer(p.HTMLURL),
			},
			Number: uint(p.GetNumber()),
			State:  githubMergeProposalStatus(p.State, p.MergedAt),
			Title:  p.GetTitle(),
			Timestamps: forgedomain.MergeProposalTimestamps{
				CreatedAt: p.GetCreatedAt().Time,
				UpdatedAt: p.GetUpdatedAt().Time,
			},
		})
	}
	return Ok(resp)
}

func (g *githubForge) GetIssues(ctx context.Context) Result[forgedomain.Issues] {
	issues, _, err := g.client.Issues.List(
		ctx,
		true,
		&github.IssueListOptions{
			Filter: "assigned",
		},
	)
	if err != nil {
		return Err[forgedomain.Issues](err)
	}
	var response forgedomain.Issues
	for _, issue := range issues {
		labels := make([]string, 0, len(issue.Labels))
		for i, l := range issue.Labels {
			labels[i] = l.GetName()
		}
		response = append(response, forgedomain.Issue{
			Author: issue.GetUser().GetName(),
			Id:     uint(issue.GetID()),
			Context: forgedomain.IssueContext{
				Assignee: extension.OptionFromPointer(issue.GetAssignee().Name),
				Labels:   labels,
				Url:      extension.OptionFromPointer(issue.HTMLURL),
			},
			Status: githubIssueStatus(issue.State),
			Title:  issue.GetTitle(),
		})
	}
	return Err[forgedomain.Issues](nil)
}

func githubMergeProposalStatus(state *string, mergedAt *github.Timestamp) forgedomain.Value[forgedomain.MergeProposalState] {
	if state == nil {
		return forgedomain.NewValue(forgedomain.MergeProposalStateNone, "")
	}
	var enum forgedomain.MergeProposalState
	switch s := *state; strings.ToUpper(s) {
	case "OPEN":
		enum = forgedomain.MergeProposalStateOpen
	case "CLOSED":
		if mergedAt != nil {
			enum = forgedomain.MergeProposalStateMerged
		} else {
			enum = forgedomain.MergeProposalStateClosed
		}
	}
	return forgedomain.NewValue(enum, state)
}

func githubNotificationSubjectTypeToKind(subjectType *string) Option[forgedomain.Value[forgedomain.NotificationKind]] {
	if subjectType == nil {
		return None[forgedomain.Value[forgedomain.NotificationKind]]()
	}
	switch v := *subjectType; strings.ToUpper(v) {
	case "ISSUE":
		return Some(forgedomain.NewValue(
			forgedomain.NotificationKindIssue,
			v,
		))
	case "PULLREQUEST":
		return Some(forgedomain.NewValue(
			forgedomain.NotificationKindMergeProposal,
			v,
		))
	default:
		return Some(forgedomain.NewValue(
			forgedomain.NotificationKindUnknown,
			v,
		))
	}
}

func githubRepository(repository *github.Repository) Option[forgedomain.Repository] {
	if repository == nil {
		return None[forgedomain.Repository]()
	}

	var group string
	if repository.GetOwner().GetName() != "" {
		group = repository.GetOwner().GetName()
	} else if repository.GetOrganization().GetName() != "" {
		group = repository.GetOrganization().GetName()
	}
	return Some(forgedomain.Repository{
		Group:   group,
		Project: repository.GetName(),
		Url:     extension.OptionFromPointer(repository.HTMLURL),
	})
}

func githubNotificationTrigger(trigger string) forgedomain.Value[forgedomain.NotificationTrigger] {
	switch strings.ToUpper(trigger) {
	case "APPROVAL_REQUESTED", "REVIEW_REQUESTED":
		return forgedomain.NewValue(forgedomain.ForgeNotificationTriggerReviewRequested, trigger)
	case "MENTION", "TEAM_MENTION":
		return forgedomain.NewValue(forgedomain.ForgeNotificationTriggerMentioned, trigger)
	case "ASSIGN":
		return forgedomain.NewValue(forgedomain.ForgeNotificationTriggerAssigned, trigger)
	default:
		return forgedomain.NewValue(forgedomain.ForgeNotificationTriggerOther, trigger)
	}
}

func githubIssueStatus(status *string) forgedomain.Value[forgedomain.IssueStatus] {
	if status == nil {
		return forgedomain.NewValue(forgedomain.IssueStatusNone, "")
	}
	var resp forgedomain.IssueStatus
	switch strings.ToUpper(*status) {
	case "OPEN":
		resp = forgedomain.IssueStatusOpen
	case "CLOSED":
		resp = forgedomain.IssueStatusClosed
	}
	return forgedomain.NewValue(resp, status)
}

func githubNotificationStatus(isRead bool) forgedomain.Value[forgedomain.NotificationStatus] {
	return forgedomain.NewValue(forgedomain.NotificationStatusPending, isRead)
}
