package manifest

import (
	"os"

	"go.yaml.in/yaml/v4"
)

type Manifest struct {
	Name         string                `yaml:"name"`
	Version      string                `yaml:"version"`
	Dependencies map[string]Dependency `yaml:"dependencies"`
}

func LoadManifestFile(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	m := &Manifest{}
	if err := yaml.Unmarshal(data, m); err != nil {
		return nil, err
	}
	if m.Dependencies == nil {
		m.Dependencies = make(map[string]Dependency)
	}
	return m, nil
}

func (m *Manifest) Write(path string) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
