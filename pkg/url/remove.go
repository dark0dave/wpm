package url

import (
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/rs/zerolog/log"
)

func Remove(m *manifest.Manifest, path, name string) error {
	_, ok := m.Dependencies[name]
	if ok {
		delete(m.Dependencies, name)
	}
	if err := m.Write(path); err != nil {
		log.Error().Msgf("Failed to write to config, %s", err)
		return err
	}
	log.Trace().Msgf("Written new config: %+v", m.Dependencies)
	return nil
}
