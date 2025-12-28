package url

import (
	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Remove(m *manifest.Manifest, path, name string) error {
	_, ok := m.Dependencies[name]
	if ok {
		delete(m.Dependencies, name)
	}
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config: %+v", slog.Any("dependencies", m.Dependencies))
	return nil
}
