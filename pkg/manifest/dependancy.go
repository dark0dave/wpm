package manifest

import (
	u "net/url"
)

type DependencyProps interface {
	Download(folderPath string) error
}

type Dependency struct {
	Name     string   `yaml:"name"`
	Url      u.URL    `yaml:"url"`
	Version  string   `yaml:"version"`
	Protocol Protocol `yaml:"protocol"`
}
