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
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)
