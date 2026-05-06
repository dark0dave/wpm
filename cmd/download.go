package cmd

import (
	"errors"
	log "log/slog"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/url"
	"github.com/dark0dave/wpm/pkg/util"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, dep := range m.Dependencies {
				log.Debug("Dependency", log.Any("dependency", dep))
				wg.Go(func() error {
					if err := download(&dep, FolderPath); err != nil {
						log.Error("Failed to download", log.Any("error", err))
						return err
					}
					return nil
				})
			}
			return wg.Wait()
		},
	}
)

func download(d *manifest.Dependency, folderPath string) error {
	switch d.Protocol {
	case manifest.Git:
		d := git.Dependency{Dependency: d}
		if err := d.Download(folderPath); err != nil {
			return err
		}
		checksum, err := util.CheckSum(folderPath)
		if err != nil {
			return err
		}
		d.CheckSum = checksum
		return nil
	case manifest.Url:
		d := url.Dependency{Dependency: d}
		if err := d.Download(folderPath); err != nil {
			return err
		}
		checksum, err := util.CheckSum(folderPath)
		if err != nil {
			return err
		}
		d.CheckSum = checksum
		return nil
	default:
		return errors.New("unsupported protocol")
	}
}
