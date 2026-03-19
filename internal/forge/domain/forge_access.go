package forgedomain

type ForgeAccess string

const (
	ForgeAccessGithubPersonalAccessToken ForgeAccess = "personal_access_token"
)

func (f ForgeAccess) String() string {
	return string(f)
}
