package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/spf13/cobra"
)

func gitAddCmd() *cobra.Command {
	var name, url, ref string
	var parsedUrl *u.URL
	cmd := &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(url)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Add(m, path, name, ref, parsedUrl)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.Flags().StringVar(&ref, "ref", "r", "")
	cmd.Flags().StringVar(&url, "url", "u", "")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("ref")

	return cmd
}

func gitRemoveCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Remove git dependencies",
		Long:    `Remove git dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Remove(m, path, name)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.MarkFlagRequired("name")

	return cmd
}

func init() {

	addCmd.AddCommand(gitAddCmd())
	rmCmd.AddCommand(gitRemoveCmd())
}
