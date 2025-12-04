package forgedomain

type NotificationKind string

const (
	NotificationKindUnknown       NotificationKind = "unknown"
	NotificationKindIssue         NotificationKind = "issue"
	NotificationKindMergeProposal NotificationKind = "merge_proposal"
)
