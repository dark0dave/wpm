package manifest

import (
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
)

var git_folder_name string = "git"

type GitDependancy struct {
	Name        string     `yaml:"name"`
	Path        string     `yaml:"path"`
	VersionType GitVersion `yaml:"version_type"`
	Version     string     `yaml:"version"`
}

func (g *GitDependancy) Download(folderPath string) (err error) {
	path := filepath.Join(folderPath, git_folder_name, g.Name)

	_, err = git.PlainClone(path, true, &git.CloneOptions{
		URL:          g.Path,
		Progress:     os.Stdout,
		SingleBranch: true,
	})
	return err
}
