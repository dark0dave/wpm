package cmd

import (
	"errors"
	"io/fs"
	"os"

	s "log/slog"

	"github.com/dark0dave/wpm/pkg/config"
	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/util"
	"github.com/spf13/cobra"
)

const (
	ManifestFileName string = "wpm"
	FolderPath       string = "weidu_modules"
)

var (
	slog                                = s.New(s.NewJSONHandler(os.Stdout, nil))
	c                                   *config.Config
	m                                   *manifest.Manifest
	path, manifestName, manifestVersion string
	rootCmd                             = &cobra.Command{
		Use:   "wpm",
		Short: "wpm is a weidu package manager",
		Long:  `A Fast and Flexible Package Manager, designed to help wiedu modders share code.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return cmd.Help()
		},
	}
)

func initConfig() {
	err := config.InitViper(slog)
	cobra.CheckErr(err)
	c, err = config.Load()
	cobra.CheckErr(err)
	m, err = manifest.LoadManifestFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		m = &manifest.Manifest{
			Name:         manifestName,
			Version:      manifestVersion,
			Dependencies: make(map[string]manifest.Dependency),
		}
		return
	}
	if err != nil {
		slog.Error("Failed to parse config file, either wpm.yaml does not exist or fails to conform to expected structure", "error", err)
		cobra.CheckErr(err)
	}
	if m.Name == "" {
		slog.Warn("Name of log empty", "path", path)
	}
	if m.Version == "" {
		slog.Warn("Version of log empty")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "wpm.yaml", "path to manifest")
	rootCmd.PersistentFlags().StringVarP(&manifestName, "manifest", "m", "New Manifest", "name for manifest")
	rootCmd.PersistentFlags().StringVarP(&manifestVersion, "x", "x", "1.0.0", "manifest version")

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	util.AddColor(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Failed", "error", err)
		os.Exit(1)
	}
}
