package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const templateText = `
Weidu Package Manager
Version:    {{.Version}}
Go version: {{.GoVersion}}
OS/Arch:    {{.OSArch}}
Git commit: {{.GitCommit}}
`

var (
	version string = "dev"
	gitSha  string = "dev"
)

func infoMessage() (string, error) {
	tmpl := template.Must(template.New("info").Parse(templateText))

	data := struct {
		Version   string
		GoVersion string
		OSArch    string
		GitCommit string
	}{
		Version:   version,
		GoVersion: runtime.Version(),
		OSArch:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GitCommit: gitSha,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Version info",
	Long:    `Detailed information about the version`,
	Run: func(cmd *cobra.Command, args []string) {
		if msg, err := infoMessage(); err != nil {
			log.Error().Msgf("Failed to print message: %s", err)
		} else {
			log.Info().Msg(msg)
		}

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
