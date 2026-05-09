package url

import (
	"fmt"
	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Add(m *manifest.Manifest, path, name, version, url string) error {
	dependency := New(name, version, url)
	if _, ok := m.Dependencies[name]; ok {
		return fmt.Errorf("Git dependency already exists: %+v", dependency)
	}
	m.Dependencies[name] = dependency.Dependency
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config", slog.Any("dependency", m.Dependencies))
	return nil
}
