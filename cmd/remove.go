package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"r", "rm"},
		Short:   "Add dependencies",
		Long:    `Add dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			// https://github.com/spf13/viper?tab=readme-ov-file#writing-config-files
			// TODO: Should be able to just write the current manifest to disk after the dep is removed ofc
			log.Debug().Msg("Added!")
		},
	}
)

func init() {
	addCmd.AddCommand(rmCmd)
}
