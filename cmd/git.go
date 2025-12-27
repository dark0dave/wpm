package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/spf13/cobra"
)

var (
	gitName, gitUrl, gitRef string
	gitAddCmd               = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			parsedUrl, err = u.Parse(gitUrl)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Add(m, path, gitName, gitRef, parsedUrl)
		},
	}
	gitRemoveCmd = &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Remove git dependencies",
		Long:    `Remove git dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return git.Remove(m, path, gitName)
		},
	}
)

func init() {
	gitAddCmd.Flags().StringVar(&gitName, "name", "n", "")
	gitAddCmd.Flags().StringVar(&gitRef, "ref", "r", "")
	gitAddCmd.Flags().StringVar(&gitUrl, "url", "u", "")
	gitAddCmd.MarkFlagRequired("name")
	gitAddCmd.MarkFlagRequired("url")
	gitAddCmd.MarkFlagRequired("ref")
	gitRemoveCmd.Flags().StringVar(&gitName, "name", "n", "")
	gitRemoveCmd.Flags().StringVar(&gitUrl, "url", "u", "")
	gitRemoveCmd.Flags().StringVar(&gitRef, "ref", "r", "")
	gitRemoveCmd.MarkFlagRequired("name")
	gitRemoveCmd.MarkFlagRequired("url")
	gitRemoveCmd.MarkFlagRequired("ref")
	addCmd.AddCommand(gitAddCmd)
	rmCmd.AddCommand(gitRemoveCmd)
}
