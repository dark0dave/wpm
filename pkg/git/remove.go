package git

import (
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/rs/zerolog/log"
)

func Remove(m *manifest.Manifest, path, name string) error {
	val, ok := m.Dependencies[name]
	if ok {
		delete(m.Dependencies, name)
	}
	log.Debug().Msgf("Removed git dependency: %+v", val)
	if err := m.Write(path); err != nil {
		log.Error().Msgf("Failed to write to config, %s", err)
		return err
	}
	log.Trace().Msg("Written new config")
	return nil
}
