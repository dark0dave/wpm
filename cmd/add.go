package cmd

import (
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Add dependencies",
		Long:    `Add dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
