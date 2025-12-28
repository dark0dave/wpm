package dropbox

import (
	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Remove(m *manifest.Manifest, path, name string) error {
	val, ok := m.Dependencies[name]
	if ok {
		delete(m.Dependencies, name)
	}
	slog.Debug("Removed git dependency", slog.Any("dependency", val))
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config, %s", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config")
	return nil
}
