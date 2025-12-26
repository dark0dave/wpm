package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

const (
	ManifestFileName string = "wpm"
	FolderPath       string = "weidu_modules"
)

var (
	m       *manifest.Manifest
	path    string
	workers int
	rootCmd = &cobra.Command{
		Use:   "wpm",
		Short: "wpm is a weidu package manager",
		Long:  `A Fast and Flexible Package Manager, designed to help wiedu modders share code.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := loadManifestFile(); err != nil {
				return err
			}
			return cmd.Help()
		},
	}
)

func loadManifestFile() error {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := yaml.Unmarshal(data, m); err != nil {
		log.Error().Msgf("Failed to parse config file, either wpm.yaml does not exist or fails to conform to expect structure: %+v", err)
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
	).Replace(rootCmd.UsageTemplate())
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{CyanStyleHeading}}`)
	rootCmd.SetUsageTemplate(usageTemplate)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "wpm.yaml", "path to manifest")
	rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 5, "number of workers for downloading mods")

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	colorize()
	if err := rootCmd.Execute(); err != nil {
		log.Error().Msgf("Failed with: %+v", err)
		os.Exit(1)
	}
}
