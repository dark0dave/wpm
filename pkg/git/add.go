package git

import (
	"fmt"
	"net/url"

	"log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
)

func Add(m *manifest.Manifest, path, name, ref string, url *url.URL) error {
	dependency := New(name, ref, *url)
	_, ok := m.Dependencies[name]
	if ok {
		return fmt.Errorf("Git dependency already exists: %#v", *dependency)
	}
	m.Dependencies[name] = *dependency.Dependency
	slog.Debug("Added git dependency", slog.Any("dependency", dependency))
	if err := m.Write(path); err != nil {
		slog.Error("Failed to write to config", slog.Any("error", err))
		return err
	}
	slog.Debug("Written new config")
	return nil
}
