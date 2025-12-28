package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/url"
	"github.com/spf13/cobra"
)

func urlAddCmd() *cobra.Command {
	var name, urlString, version string
	var parsedUrl *u.URL
	cmd := &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Add url dependencies",
		Long:    `Add url dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(urlString)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return url.Add(m, path, name, version, parsedUrl)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.Flags().StringVar(&urlString, "url", "u", "")
	cmd.Flags().StringVar(&version, "version", "v", "")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("version")

	return cmd
}

func urlRemoveCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Remove url dependencies",
		Long:    `Remove url dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return url.Remove(m, path, name)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.MarkFlagRequired("name")

	return cmd
}

func init() {
	addCmd.AddCommand(urlAddCmd())
	rmCmd.AddCommand(urlRemoveCmd())
}
