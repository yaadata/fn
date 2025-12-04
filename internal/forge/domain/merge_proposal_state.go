package forgedomain

type MergeProposalState string

const (
	MergeProposalStateNone   MergeProposalState = ""
	MergeProposalStateOpen   MergeProposalState = "open"
	MergeProposalStateClosed MergeProposalState = "closed"
	MergeProposalStateMerged MergeProposalState = "merged"
)

func (p MergeProposalState) String() string {
	return string(p)
}
