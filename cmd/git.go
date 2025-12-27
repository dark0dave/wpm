package cmd

import (
	"fmt"
	u "net/url"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	gitName, gitUrl, gitRef string
	gitAddCmd               = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(gitUrl)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dependency := git.New(gitName, gitRef, *parsedUrl)
			_, ok := m.Dependencies[gitName]
			if ok {
				return fmt.Errorf("Git dependency already exists: %+v", dependency)
			}
			m.Dependencies[gitName] = *dependency.Dependency
			log.Debug().Msgf("Added git dependency: %+v", dependency)
			if err := m.Write(path); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
				return err
			}
			log.Trace().Msg("Written new config")
			return nil
		},
	}
	gitRemoveCmd = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Remove git dependencies",
		Long:    `Remove git dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			val, ok := m.Dependencies[gitName]
			if ok {
				delete(m.Dependencies, gitName)
			}
			log.Debug().Msgf("Removed git dependency: %+v", val)
			if err := m.Write(path); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
				return err
			}
			log.Trace().Msg("Written new config")
			return nil
		},
	}
)

func init() {
	gitAddCmd.Flags().StringVar(&gitName, "name", "n", "")
	gitAddCmd.Flags().StringVar(&gitRef, "ref", "r", "")
	gitAddCmd.Flags().StringVar(&gitUrl, "url", "u", "")
	gitAddCmd.MarkFlagRequired("name")
	gitAddCmd.MarkFlagRequired("url")
	gitAddCmd.MarkFlagRequired("ref")
	gitRemoveCmd.Flags().StringVar(&gitName, "name", "n", "")
	gitRemoveCmd.Flags().StringVar(&gitUrl, "url", "u", "")
	gitRemoveCmd.Flags().StringVar(&gitRef, "ref", "r", "")
	gitRemoveCmd.MarkFlagRequired("name")
	gitRemoveCmd.MarkFlagRequired("url")
	gitRemoveCmd.MarkFlagRequired("ref")
	addCmd.AddCommand(gitAddCmd)
	rmCmd.AddCommand(gitRemoveCmd)
}
