package url

import (
	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Remove(m *manifest.Manifest, path, name string) error {
	val, ok := m.Dependencies[name]
	if !ok {
		slog.Debug("Did not find url dependency", slog.String("name", name))
		return nil
	}
	delete(m.Dependencies, name)
	slog.Debug("Removed git dependency", slog.Any("dependency", val))
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config: %+v", slog.Any("dependencies", m.Dependencies))
	return nil
}
