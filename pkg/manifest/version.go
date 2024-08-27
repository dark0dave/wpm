package manifest

import (
	"errors"
)

type GitVersion string

func (g GitVersion) Set(s string) error {
	if s != string(branch) && s != string(tag) {
		return errors.New("Is not one of the following: rev, branch or tag")
	}
	g = GitVersion(s)
	return nil
}

func (g GitVersion) String() string {
	return string(g)
}

func (g GitVersion) Type() string {
	return "GitVersion one of rev, branch or tag"
}

const (
	branch GitVersion = "branch"
	tag    GitVersion = "tag"
)
