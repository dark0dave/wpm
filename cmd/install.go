package cmd

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	wg         sync.WaitGroup
	installCmd = &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install all the dependencies from your project file (wpm.yaml)",
		Long: `Install all the dependencies in your project
from wpm.yaml file to the weidu_modules folder`,
		Run: func(cmd *cobra.Command, args []string) {
			wg.Add(workers)
			for _, dep := range m.Dependencies {
				log.Debug().Msgf("Dep: %+v\n", dep)
				go func() {
					defer wg.Done()
					if err := dep.Download(FolderPath); err != nil {
						log.Error().Msgf("Failed to install, %s", err)
					}
				}()
			}
			wg.Wait()
		},
	}
)
