package cmd

import (
	"errors"
	"os"
	"path/filepath"

	s "log/slog"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/dark0dave/wpm/pkg/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ManifestFileName string = "wpm"
	FolderPath       string = "weidu_modules"
)

var (
	slog    = s.New(s.NewJSONHandler(os.Stdout, nil))
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

func initConfig() {
	configPath := os.Getenv("XDGHOME")
	if configPath == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		configPath = filepath.Join(home, ".config")
	}

	filePath := filepath.Join(configPath, "wpm", "default.yaml")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	viper.WithLogger(slog)

	err := viper.ReadInConfig()
	if err == nil {
		slog.Debug("Using config file: %s", "file", viper.ConfigFileUsed())
		return
	}

	if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		slog.Error("Config fail to initialize", "error", err)
		return
	}

	slog.Info("First time run, creating the config file at", "directory", filePath)

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		slog.Error("Could not create the config directory", "error", err)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		slog.Error("Could not create the config file", "error", err)
		return
	}
	if err := viper.WriteConfigTo(f); err != nil {
		slog.Error("Could not create the config file", "error", err)
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "wpm.yaml", "path to manifest")

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	util.AddColor(rootCmd)
	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Debug("Config file changed: %+v", e)
	})
	viper.WatchConfig()
	if err := rootCmd.Execute(); err != nil {
		slog.Debug("Failed with: %+v", err)
		os.Exit(1)
	}
}
