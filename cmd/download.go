package cmd

import (
	log "log/slog"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	wg          errgroup.Group
	downloadCmd = &cobra.Command{
		Use:     "download",
		Aliases: []string{"d"},
		Short:   "Download all the dependencies from your project file (wpm.yaml)",
		Long: `Download all the dependencies in your project
from wpm.yaml file to the weidu_modules folder`,
		RunE: func(_ *cobra.Command, _ []string) error {
			for _, dep := range m.Dependencies {
				log.Debug("Dependency", log.Any("dependency", dep))
				wg.Go(func() error {
					if err := dep.Download(FolderPath); err != nil {
						log.Error("Failed to download", log.Any("error", err))
						return err
					}
					if err := dep.CheckSum(FolderPath); err != nil {
						log.Error("Failed to create valid checksum", log.Any("error", err))
						return err
					}
					return nil
				})
			}
			return wg.Wait()
		},
	}
)
