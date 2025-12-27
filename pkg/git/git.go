package git

import (
	"os"
	"path/filepath"

	u "net/url"

	"github.com/dark0dave/wpm/pkg/manifest"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Dependency struct {
	*manifest.Dependency
}

func New(name, ref string, url u.URL) *Dependency {
	out := &Dependency{
		Dependency: &manifest.Dependency{
			Name:     name,
			Url:      url,
			Version:  plumbing.ReferenceName(ref).String(),
			Protocol: manifest.Git,
		},
	}
	return out
}

func (g *Dependency) Download(folderPath string) error {
	path := filepath.Join(folderPath, g.Name)
	_, err := git.PlainClone(path, true, &git.CloneOptions{
		URL:           g.Url.String(),
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName(g.Version),
		SingleBranch:  true,
		Depth:         0,
		Tags:          3,
	})
	return err
}
