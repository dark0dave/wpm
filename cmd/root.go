package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ManifestFileName string = "wpm"
const FolderPath string = "weidu_modules"

var (
	m       manifest.Manifest
	workers int = 2
	rootCmd     = &cobra.Command{
		Use:   "wpm",
		Short: "wpm is a weidu package manager",
		Long:  `A Fast and Flexible Package Manager, designed to help wiedu modders share code.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func loadManifestFile() (err error) {
	viper.SetConfigName(ManifestFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&m); err != nil {
		log.Error().AnErr("Failed to parse config file, either wpm.yaml does not exist or fails to conform to expect structure", err)
		return err
	}
	return nil
}

func colorize() {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	cobra.AddTemplateFunc("CyanStyleHeading", color.New(color.BgCyan).SprintFunc())
	usageTemplate := strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
		// The following one steps on "Global Flags:"
		// `Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(rootCmd.UsageTemplate())
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{CyanStyleHeading}}`)
	rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", workers, "Number of works for wpm")
	rootCmd.SetUsageTemplate(usageTemplate)
}

func Execute() {
	colorize()
	if err := loadManifestFile(); err != nil {
		log.Error().AnErr("Failed with", err)
		os.Exit(1)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Debug().Msgf("Config file changed: %+v", e)
	})
	viper.WatchConfig()
	if err := rootCmd.Execute(); err != nil {
		log.Error().AnErr("Failed", err)
		os.Exit(1)
	}
}
