package manifest

type Manifest struct {
	Name         string       `yaml:"name"`
	Version      string       `yaml:"version"`
	Dependancies Dependancies `mapstructure:"dependancies"`
}

type Dependancies struct {
	GitDependancies []GitDependancy `mapstructure:"git"`
	UrlDependancies []UrlDependancy `mapstructure:"url"`
}

type Dependancy interface {
	GetName() string
	Download(folderPath string) error
}
