package cmd

import (
	"errors"
	"io/fs"
	s "log/slog"
	"os"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/util"
	"github.com/spf13/cobra"
)

const (
	ManifestFileName string = "wpm"
	FolderPath       string = "weidu_modules"
	DefaultVersion   string = "1.0.0"
	DefaultName      string = "new"
)

var (
	slog                                = s.New(s.NewJSONHandler(os.Stdout, nil))
	m                                   *manifest.Manifest
	path, manifestName, manifestVersion string
	rootCmd                             = &cobra.Command{
		Use:   "wpm",
		Short: "wpm is a weidu package manager",
		Long:  `A Fast and Flexible Package Manager, designed to help wiedu modders share code.`,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			if m.Version == "" || manifestVersion != DefaultVersion {
				m.Version = manifestVersion
			}
			if m.Name == "" || manifestVersion != DefaultName {
				m.Name = manifestName
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			return cmd.Help()
		},
	}
)

func initConfig() {
	var err error
	m, err = manifest.LoadManifestFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		m = &manifest.Manifest{
			Dependencies: make(map[string]manifest.Dependency),
		}
		return
	}
	if err != nil {
		slog.Error("Failed to parse config file, either wpm.yaml does not exist or fails to conform to expected structure", "error", err)
		cobra.CheckErr(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "wpm.yaml", "path to manifest")
	rootCmd.PersistentFlags().StringVarP(&manifestName, "manifest", "m", DefaultName, "name for manifest")
	rootCmd.PersistentFlags().StringVarP(&manifestVersion, "x", "x", DefaultVersion, "manifest version")

	rootCmd.AddCommand(downloadCmd, addCmd, rmCmd, versionCmd, logCmd())
}

func Execute() {
	util.AddColor(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Failed", s.Any("error", err))
		os.Exit(1)
	}
}
