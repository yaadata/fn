package forgedomain

type MergeProposal struct {
	Author     string
	Context    MergeProposalContext
	Repository Repository
	Number     uint
	State      Value[MergeProposalState]
	Title      string
	Timestamps MergeProposalTimestamps
}
