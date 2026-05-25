package cmd

import (
	"github.com/spf13/cobra"
)

func logCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "log",
		Aliases: []string{"lg"},
		Short:   "Create a weidu log file",
		Long:    `Create a weidu log file from the configuration provided.`,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}

	return cmd
}

func init() {
	rmCmd.AddCommand(urlRemoveCmd())
}
