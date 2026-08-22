package cmd

import (
	"errors"
	u "net/url"
	"strings"

	"github.com/dark0dave/wpm/pkg/git"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

var ErrorName = errors.New("could not construct a name for the mod, try again with --name")

func gitAddCmd() *cobra.Command {
	var name, url, ref string
	var version plumbing.ReferenceName
	cmd := &cobra.Command{
		Use:     "git",
		Aliases: []string{"g"},
		Short:   "Add git dependencies",
		Long:    `Add git dependencies to a manifest file`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			version = plumbing.ReferenceName(ref)
			if err := version.Validate(); err != nil {
				version = plumbing.HEAD
				slog.Debug("defaulting to head")
			}
			if err := cmd.MarkFlagRequired("url"); err != nil {
				return err
			}
			u, err := u.Parse(url)
			if err != nil {
				return err
			}
			if name != "" {
				return nil
			}
			if p := strings.Split(u.Path, "/"); len(p) > 2 {
				name = strings.ToLower(p[2])
				return nil
			}
			return ErrorName
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return git.Add(m, path, name, version.Short(), url)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "")
	cmd.Flags().StringVar(&ref, "ref", "", "")
	cmd.Flags().StringVar(&url, "url", "", "")

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
