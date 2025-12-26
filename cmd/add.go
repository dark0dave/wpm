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
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)
