package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/dropbox"
	"github.com/spf13/cobra"
)

func dropboxAddCmd() *cobra.Command {
	var name, version, url string
	var parsedUrl *u.URL
	cmd := &cobra.Command{
		Use:     "dropbox",
		Aliases: []string{"d"},
		Short:   "Add dropbox dependencies",
		Long:    `Add dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(url)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return dropbox.Add(m, path, name, version, parsedUrl)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.Flags().StringVar(&url, "url", "u", "")
	cmd.Flags().StringVar(&version, "version", "v", "")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("version")

	return cmd
}

func dropboxRemoveCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:     "dropbox",
		Aliases: []string{"d"},
		Short:   "Remove dropbox dependencies",
		Long:    `Remove dependencies from the manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dropbox.Remove(m, path, name)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.MarkFlagRequired("name")

	return cmd
}

func init() {
	addCmd.AddCommand(dropboxAddCmd())
	rmCmd.AddCommand(dropboxRemoveCmd())
}
