package forgedomain

type Provider string

const (
	ForgeProviderGitHub   Provider = "github"
	ForgeProviderCodeberg Provider = "codeberg"
)

func (p Provider) String() string {
	return string(p)
}
