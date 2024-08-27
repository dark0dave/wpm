package manifest

import "errors"

type GitVersion string

func (v GitVersion) Set(s string) error {
	if s != string(rev) && s != string(branch) && s != string(tag) {
		return errors.New("Is not one of the following: rev, branch or tag")
	}
	v = GitVersion(s)
	return nil
}

func (v GitVersion) String() string {
	return string(v)
}

func (v GitVersion) Type() string {
	return "GitVersion one of rev, branch or tag"
}

const (
	rev    GitVersion = "rev"
	branch GitVersion = "branch"
	tag    GitVersion = "tag"
)
