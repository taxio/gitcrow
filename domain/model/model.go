package model

type GitRepo struct {
	Owner   string
	Repo    string
	Tag     string
	IsClone bool
}

type Report struct {
	GitRepo *GitRepo
	Success bool
	Message string
}

var (
	LatestTag = "latest:tag" // Since colons cannot be used in git tag names
)
