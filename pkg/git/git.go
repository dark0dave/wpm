package git

import (
	"os"
	"path/filepath"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/util"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Dependency struct {
	*manifest.Dependency
}

func New(name, version, url string) *Dependency {
	return &Dependency{
		Dependency: &manifest.Dependency{
			Name:     name,
			URL:      url,
			Version:  version,
			Protocol: manifest.Git.String(),
		},
	}
}

func (g *Dependency) Download(folderPath string) error {
	path := filepath.Join(folderPath, g.Name)
	_, err := git.PlainClone(path, true, &git.CloneOptions{
		URL:           g.URL,
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName(g.Version),
		SingleBranch:  true,
		Depth:         0,
		Tags:          3,
	})
	return err
}

func (g *Dependency) CheckSum(folderPath string) error {
	checksum, err := util.CheckSum(folderPath)
	if err != nil {
		return err
	}
	g.Dependency.CheckSum = checksum
	return nil
}
