package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/url"
	"github.com/spf13/cobra"
)

func urlAddCmd() *cobra.Command {
	var name, urlString, version string
	cmd := &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Add url dependencies",
		Long:    `Add url dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := cmd.MarkFlagRequired("name"); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired("url"); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired("version"); err != nil {
				return err
			}
			_, err := u.Parse(urlString)
			return err
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return url.Add(m, path, name, version, urlString)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.Flags().StringVar(&urlString, "url", "u", "")
	cmd.Flags().StringVar(&version, "version", "v", "")

	return cmd
}

func urlRemoveCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Remove url dependencies",
		Long:    `Remove url dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.MarkFlagRequired("name")
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return url.Remove(m, path, name)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")

	return cmd
}

func init() {
	addCmd.AddCommand(urlAddCmd())
	rmCmd.AddCommand(urlRemoveCmd())
}
