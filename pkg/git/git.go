package git

import (
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Dependency struct {
	Name string                 `yaml:"name"`
	Url  string                 `yaml:"url"`
	Ref  plumbing.ReferenceName `yaml:"ref"`
}

func New(name, url, ref string) *Dependency {
	return &Dependency{
		Name: name,
		Url:  url,
		Ref:  plumbing.ReferenceName(ref),
	}
}

func (g *Dependency) Download(folderPath string) error {
	path := filepath.Join(folderPath, g.Name)
	_, err := git.PlainClone(path, true, &git.CloneOptions{
		URL:           g.Url,
		Progress:      os.Stdout,
		ReferenceName: g.Ref,
		SingleBranch:  true,
		Depth:         0,
		Tags:          3,
	})
	return err
}
