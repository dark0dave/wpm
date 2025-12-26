package manifest

import (
	"os"

	"github.com/rs/zerolog/log"
	"go.yaml.in/yaml/v3"
)

type Manifest struct {
	Name         string       `yaml:"name"`
	Version      string       `yaml:"version"`
	Dependencies []Dependency `mapstructure:"dependencies"`
}

func LoadManifestFile(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	m := &Manifest{}
	if err := yaml.Unmarshal(data, m); err != nil {
		log.Error().Msgf("Failed to parse config file, either wpm.yaml does not exist or fails to conform to expect structure: %+v", err)
		return nil, err
	}
	return m, nil
}
