package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/url"
	"github.com/spf13/cobra"
)

var (
	urlName, urlString, urlVersion string
	parsedUrl                      *u.URL
	urlAddCmd                      = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Add url dependencies",
		Long:    `Add url dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(urlString)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return url.Add(m, path, urlName, urlVersion, parsedUrl)
		},
	}
	urlRemoveCmd = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Remove url dependencies",
		Long:    `Remove url dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return url.Remove(m, path, urlName)
		},
	}
)

func init() {
	urlAddCmd.Flags().StringVar(&urlName, "name", "n", "")
	urlAddCmd.Flags().StringVar(&urlString, "url", "u", "")
	urlAddCmd.Flags().StringVar(&urlVersion, "version", "v", "")
	urlAddCmd.MarkFlagRequired("name")
	urlAddCmd.MarkFlagRequired("url")
	urlAddCmd.MarkFlagRequired("version")
	urlRemoveCmd.Flags().StringVar(&urlName, "name", "n", "")
	urlRemoveCmd.Flags().StringVar(&urlString, "url", "u", "")
	urlRemoveCmd.Flags().StringVar(&urlVersion, "version", "v", "")
	urlRemoveCmd.MarkFlagRequired("name")
	urlRemoveCmd.MarkFlagRequired("url")
	urlRemoveCmd.MarkFlagRequired("version")
	addCmd.AddCommand(urlAddCmd)
	rmCmd.AddCommand(urlRemoveCmd)
}
