package cmd

import (
	"github.com/dark0dave/wpm/pkg/git"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitName, gitUrl, gitRef string
	gitAddCmd               = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			gitDependency := git.New(gitName, gitUrl, gitRef)
			for _, dep := range m.Dependencies {
				if dep == gitDependency {
					log.Error().Msgf("Git dependency already exists: %+v", dep)
					cmd.ErrOrStderr().Write([]byte("Failed"))
					return
				}
			}
			viper.Set("dependencies.git", append(m.Dependencies, gitDependency))
			log.Debug().Msgf("Added git dependency: %+v", viper.Get("dependencies.git"))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
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
			gitDependency := git.New(gitName, gitUrl, gitRef)
			for i, dep := range m.Dependencies {
				if dep == gitDependency {
					viper.Set("dependencies.git", append(m.Dependencies[:i], m.Dependencies[i+1:]...))
					log.Debug().Msgf("Removed git dependency: %+v", dep)
					break
				}
			}
			log.Debug().Msgf("Removed git dependency: %+v", viper.Get("dependencies.git"))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
			}
			log.Debug().Msg("Written new config")
		},
	}
)

func init() {
	gitAddCmd.PersistentFlags().StringVarP(&gitName, "name", "n", "", "")
	gitAddCmd.PersistentFlags().StringVarP(&gitUrl, "path", "p", "", "")
	gitAddCmd.PersistentFlags().StringVarP(&gitRef, "ref", "r", "", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitName, "name", "n", "", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitUrl, "path", "p", "", "")
	gitrmCmd.PersistentFlags().StringVarP(&gitRef, "ref", "r", "", "")
	addCmd.AddCommand(gitAddCmd)
	rmCmd.AddCommand(gitrmCmd)
}
