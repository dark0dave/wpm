package cmd

import (
	"errors"
	log "log/slog"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/url"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	wg            errgroup.Group
	errorProtocol = errors.New("invalid protocol")
	downloadCmd   = &cobra.Command{
		Use:     "download",
		Aliases: []string{"d"},
		Short:   "Download all the dependencies from your project file (wpm.yaml)",
		Long: `Download all the dependencies in your project
from wpm.yaml file to the weidu_modules folder`,
		RunE: func(_ *cobra.Command, _ []string) error {
			for _, dep := range m.Dependencies {
				switch manifest.Parse(dep.Protocol) {
				case manifest.Git:
					download(&git.Dependency{Dependency: &dep})
				case manifest.URL:
					download(&url.Dependency{Dependency: &dep})
				default:
					return errorProtocol
				}
			}
			return wg.Wait()
		},
	}
)

func download(d manifest.DependencyProps) {
	log.Debug("Dependency", log.Any("dependency", d))
	wg.Go(func() error {
		if err := d.Download(FolderPath); err != nil {
			log.Error("Failed to download", log.Any("error", err))
			return err
		}
		if err := d.CheckSum(FolderPath); err != nil {
			log.Error("Failed to create valid checksum", log.Any("error", err))
			return err
		}
		return nil
	})
}
