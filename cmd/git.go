package cmd

import (
	u "net/url"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/spf13/cobra"
)

func gitAddCmd() *cobra.Command {
	var name, url, ref string
	cmd := &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := cmd.MarkFlagRequired("name"); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired("url"); err != nil {
				return err
			}
			if err := cmd.MarkFlagRequired("ref"); err != nil {
				return err
			}
			_, err := u.Parse(url)
			return err
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return git.Add(m, path, name, ref, url)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")
	cmd.Flags().StringVar(&ref, "ref", "r", "")
	cmd.Flags().StringVar(&url, "url", "u", "")

	return cmd
}

func gitRemoveCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Remove git dependencies",
		Long:    `Remove git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.MarkFlagRequired("name")
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return git.Remove(m, path, name)
		},
	}

	cmd.Flags().StringVar(&name, "name", "n", "")

	return cmd
}

func init() {
	addCmd.AddCommand(gitAddCmd())
	rmCmd.AddCommand(gitRemoveCmd())
}
