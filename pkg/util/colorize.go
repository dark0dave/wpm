package util

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func AddColor(cmd *cobra.Command) {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgGreen).SprintFunc())
	cobra.AddTemplateFunc("CyanStyleHeading", color.New(color.BgCyan).SprintFunc())
	usageTemplate := strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Aliases:`, `{{StyleHeading "Aliases:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
	).Replace(cmd.UsageTemplate())
	re := regexp.MustCompile(`(?m)^Flags:\s*$`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{StyleHeading "Flags:"}}`)
	usageTemplate = re.ReplaceAllLiteralString(usageTemplate, `{{CyanStyleHeading}}`)
	cmd.SetUsageTemplate(usageTemplate)
}
