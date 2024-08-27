package manifest

type Manifest struct {
	Name         string       `yaml:"name"`
	Version      string       `yaml:"version"`
	Dependencies Dependencies `mapstructure:"dependencies"`
}

type Dependencies struct {
	GitDependencies []GitDependancy `mapstructure:"git"`
	UrlDependencies []UrlDependancy `mapstructure:"url"`
}
