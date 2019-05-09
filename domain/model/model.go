package model

import "fmt"

type GitRepo struct {
	Owner   string
	Repo    string
	Tag     string
	IsClone bool
}

type Report struct {
	GitRepo *GitRepo
	Code    ReportCode
	Message string
}

type ReportCode int

func (r ReportCode) ToString() string {
	return fmt.Sprintf("%d", r)
}

const (
	ReportSuccess ReportCode = iota
	ReportTagNotFound
	ReportInternalErr
	ReportSaveErr
)
