package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Add dependencies",
		Long:    `Add dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debug().Msg("Added!")
		},
	}
)

func init() {
	addCmd.AddCommand(installCmd)
}
