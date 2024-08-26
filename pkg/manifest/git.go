package manifest

import (
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
)

var git_folder_name string = "git"

type Version string

const (
	rev    Version = "rev"
	branch Version = "branch"
	Tag    Version = "tag"
)

type GitDependancy struct {
	Name        string  `yaml:"name"`
	Path        string  `yaml:"path"`
	VersionType Version `yaml:"version_type"`
	Version     string  `yaml:"version"`
}

func (g *GitDependancy) GetName() string {
	return g.Name
}

func (g *GitDependancy) Download(folderPath string) (err error) {
	path := filepath.Join(folderPath, git_folder_name, g.Name)

	if _, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      g.Path,
		Progress: os.Stdout,
	}); err != nil {
		return err
	}

	return nil
}
