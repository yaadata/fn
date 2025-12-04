package forgedomain

type MergeProposal struct {
	Author     string
	Context    MergeProposalContext
	Number     uint
	State      Value[MergeProposalState]
	Title      string
	Timestamps MergeProposalTimestamps
}
