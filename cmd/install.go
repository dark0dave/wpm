package cmd

import (
	"errors"

	"github.com/dark0dave/wpm/pkg/dropbox"
	"github.com/dark0dave/wpm/pkg/git"
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/url"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	wg         errgroup.Group
	installCmd = &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install all the dependencies from your project file (wpm.yaml)",
		Long: `Install all the dependencies in your project
from wpm.yaml file to the weidu_modules folder`,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, dep := range m.Dependencies {
				log.Debug().Msgf("Dep: %+v\n", dep)
				wg.Go(func() error {
					if err := download(&dep, FolderPath); err != nil {
						log.Error().Msgf("Failed to install, %s", err)
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
	case manifest.DropBox:
		d := dropbox.Dependency{Dependency: d}
		return d.Download(folderPath)
	case manifest.Git:
		d := git.Dependency{Dependency: d}
		return d.Download(folderPath)
	case manifest.Url:
		d := url.Dependency{Dependency: d}
		return d.Download(folderPath)
	default:
		return errors.New("unsupported protocol")
	}
}
