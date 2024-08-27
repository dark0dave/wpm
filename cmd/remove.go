package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rmCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"r", "rm"},
		Short:   "Removes dependencies",
		Long:    `Remove dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(rmCmd)
}
