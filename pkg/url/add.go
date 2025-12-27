package url

import (
	"fmt"
	"net/url"

	"github.com/dark0dave/wpm/pkg/manifest"

	"github.com/rs/zerolog/log"
)

func Add(m *manifest.Manifest, path, name, version string, url *url.URL) error {
	dependency := New(name, version, *url)
	_, ok := m.Dependencies[name]
	if ok {
		return fmt.Errorf("Git dependency already exists: %+v", dependency)
	}
	m.Dependencies[name] = *dependency.Dependency
	if err := m.Write(path); err != nil {
		log.Error().Msgf("Failed to write to config, %s", err)
		return err
	}
	log.Trace().Msgf("Written new config: %+v", m.Dependencies)
	return nil
}
