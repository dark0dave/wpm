package git

import (
	"fmt"
	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Add(m *manifest.Manifest, path, name, ref, url string) error {
	if dependency, ok := m.Dependencies[name]; ok {
		return fmt.Errorf("Git dependency already exists: %#v", dependency)
	}
	dependency := New(name, ref, url)

	m.Dependencies[name] = dependency.Dependency
	slog.Debug("Added git dependency", slog.Any("dependency", dependency))
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config")
	return nil
}
