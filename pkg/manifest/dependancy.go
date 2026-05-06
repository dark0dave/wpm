package manifest

type DependencyProps interface {
	Download(folderPath string) error
}

type Dependency struct {
	Name     string   `yaml:"name"`
	Url      string   `yaml:"url"`
	Version  string   `yaml:"version"`
	CheckSum string   `yaml:"checksum"`
	Protocol Protocol `yaml:"protocol"`
}
