package git

import (
	"fmt"
	"net/url"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/rs/zerolog/log"
)

func Add(m *manifest.Manifest, path, name, ref string, url *url.URL) error {
	dependency := New(name, ref, *url)
	_, ok := m.Dependencies[name]
	if ok {
		return fmt.Errorf("Git dependency already exists: %#v", *dependency)
	}
	m.Dependencies[name] = *dependency.Dependency
	log.Debug().Msgf("Added git dependency: %+v", dependency)
	if err := m.Write(path); err != nil {
		log.Error().Msgf("Failed to write to config, %s", err)
		return err
	}
	log.Trace().Msg("Written new config")
	return nil
}
