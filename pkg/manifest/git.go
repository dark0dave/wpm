package manifest

import (
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var git_folder_name string = "git"

type GitDependancy struct {
	Name        string     `yaml:"name"`
	Path        string     `yaml:"path"`
	VersionType GitVersion `yaml:"version_type"`
	Version     string     `yaml:"version"`
}

func (g GitDependancy) ToReferenceName() plumbing.ReferenceName {
	switch g.VersionType {
	case branch:
		return plumbing.NewBranchReferenceName(g.Version)
	case tag:
		return plumbing.NewTagReferenceName(g.Version)
	default:
		return plumbing.NewRemoteHEADReferenceName("origin")
	}
}

func (g *GitDependancy) Download(folderPath string) (err error) {
	path := filepath.Join(folderPath, git_folder_name, g.Name)
	referenceName := g.ToReferenceName()
	_, err = git.PlainClone(path, true, &git.CloneOptions{
		URL:           g.Path,
		Progress:      os.Stdout,
		ReferenceName: referenceName,
		SingleBranch:  true,
	})
	return err
}
