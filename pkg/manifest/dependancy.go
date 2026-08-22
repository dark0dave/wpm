package manifest

type DependencyProps interface {
	Download(folderPath string) error
	CheckSum(folderPath string) error
}

type Dependency struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Version  string `yaml:"version"`
	CheckSum string `yaml:"checksum"`
	Protocol string `yaml:"protocol"`
}
