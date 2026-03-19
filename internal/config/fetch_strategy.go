package config

type FetchStrategy string

const (
	FetchStrategyPull FetchStrategy = "pull"
	FetchStrategyPush FetchStrategy = "push"
)

func (fs FetchStrategy) String() string {
	return string(fs)
}
