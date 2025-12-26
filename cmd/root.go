package cmd

import (
	"os"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	ManifestFileName string = "wpm"
	FolderPath       string = "weidu_modules"
)

var (
	m       *manifest.Manifest
	path    string
	rootCmd = &cobra.Command{
		Use:   "wpm",
		Short: "wpm is a weidu package manager",
		Long:  `A Fast and Flexible Package Manager, designed to help wiedu modders share code.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if m, err = manifest.LoadManifestFile(path); err != nil {
				return err
			}
			return cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "wpm.yaml", "path to manifest")

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	util.AddColor(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Error().Msgf("Failed with: %+v", err)
		os.Exit(1)
	}
}
