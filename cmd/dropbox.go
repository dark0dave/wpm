package cmd

import "github.com/spf13/cobra"

var (
	dropboxAddCmd = &cobra.Command{
		Use:     "dropbox",
		Aliases: []string{"d"},
		Short:   "Add dropbox dependencies",
		Long:    `Add dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	dropboxRmCmd = &cobra.Command{
		Use:     "dropbox",
		Aliases: []string{"d"},
		Short:   "Remove dropbox dependencies",
		Long:    `Add dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func init() {
	addCmd.AddCommand(dropboxAddCmd)
	rmCmd.AddCommand(dropboxRmCmd)
}
