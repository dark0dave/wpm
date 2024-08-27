package cmd

import (
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitName        string
	gitPath        string
	gitVersionType manifest.GitVersion
	gitVersion     string
	gitAddCmd      = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			gitDependency := manifest.GitDependancy{
				Name:        gitName,
				Path:        gitPath,
				VersionType: gitVersionType,
				Version:     gitVersion,
			}
			for _, dep := range m.Dependencies.GitDependencies {
				if dep == gitDependency {
					log.Error().Msgf("Git dependency already exists: %+v", dep)
					cmd.ErrOrStderr().Write([]byte("Failed"))
					return
				}
			}
			viper.Set("dependencies.git", append(m.Dependencies.GitDependencies, gitDependency))
			log.Debug().Msgf("Added git dependency: %+v", viper.Get("dependencies.git"))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().AnErr("Failed to write to config", err)
			}
			log.Debug().Msg("Written new config")
		},
	}
	gitrmCmd = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Remove git dependencies",
		Long:    `Remove git dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			gitDependency := manifest.GitDependancy{
				Name:        gitName,
				Path:        gitPath,
				VersionType: gitVersionType,
				Version:     gitVersion,
			}
			for i, dep := range m.Dependencies.GitDependencies {
				if dep == gitDependency {
					viper.Set("dependencies.git", append(m.Dependencies.GitDependencies[:i], m.Dependencies.GitDependencies[i+1:]...))
					log.Debug().Msgf("Removed git dependency: %+v", dep)
					break
				}
			}
			log.Debug().Msgf("Removed git dependency: %+v", viper.Get("dependencies.git"))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().AnErr("Failed to write to config", err)
			}
			log.Debug().Msg("Written new config")
		},
	}
)

func init() {
	gitAddCmd.PersistentFlags().StringVarP(&gitName, "name", "n", "", "")
	gitAddCmd.PersistentFlags().StringVarP(&gitPath, "path", "p", "", "")
	gitAddCmd.PersistentFlags().VarP(gitVersionType, "type", "t", "")
	gitAddCmd.PersistentFlags().StringVarP(&gitVersion, "version", "v", "", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitName, "name", "n", "", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitPath, "path", "p", "", "")
	gitrmCmd.PersistentFlags().VarP(gitVersionType, "type", "t", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitVersion, "version", "v", "", "")
	addCmd.AddCommand(gitAddCmd)
	rmCmd.AddCommand(gitrmCmd)
}
