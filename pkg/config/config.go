package config

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFile = "config.yaml"
)

type Config struct {
	*dropbox.Config
}

func resolveConfigFilePath() (string, error) {
	configPath := os.Getenv("XDG_CONFIG_HOME")
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configPath = filepath.Join(home, ".config")
	}
	configPath = filepath.Join(configPath, "wpm")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return filepath.Join(configPath, DefaultConfigFile), nil
}

func InitViper(l *slog.Logger) error {
	configFilePath, err := resolveConfigFilePath()
	if err != nil {
		return err
	}
	viper.SetConfigFile(configFilePath)
	viper.AutomaticEnv()
	viper.WithLogger(l)

	viper.OnConfigChange(func(e fsnotify.Event) {
		l.Debug("Config file changed", "event", e)
	})
	viper.WatchConfig()
	return nil
}

func Load() (*Config, error) {
	err := viper.ReadInConfig()
	if errors.Is(err, fs.ErrNotExist) {
		return &Config{}, nil
	}
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
