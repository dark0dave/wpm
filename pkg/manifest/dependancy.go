package manifest

import "net/url"

type DependencyProps interface {
	Download(folderPath string) error
}

type Dependency struct {
	Name     string   `yaml:"name"`
	Url      url.URL  `yaml:"url"`
	Version  string   `yaml:"version"`
	Protocol Protocol `yaml:"protocol"`
}

func (d *Dependency) Download(folderPath string) error {
	panic("Unimplemented")
}
