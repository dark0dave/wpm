package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install all the dependencies in your project",
		Long: `Install all the dependencies in your project
from wpm.yaml file to the weidu_modules folder`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, dep := range m.Dependancies.GitDependancies {
				log.Debug().Msgf("Git Dep: %+v\n", dep)
				if err := dep.Download(FolderPath); err != nil {
					log.Error().AnErr("Failed to install", err)
				}
			}
			for _, dep := range m.Dependancies.UrlDependancies {
				log.Debug().Msgf("Url Dep: %+v\n", dep)
				if err := dep.Download(FolderPath); err != nil {
					log.Error().AnErr("Failed to install", err)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}