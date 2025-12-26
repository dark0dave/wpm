package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const templateText = `
Weidu Package Manager
Version:    {{.Version}}
Go version: {{.GoVersion}}
OS/Arch:    {{.OSArch}}
`

func infoMessage() (*string, error) {
	tmpl, err := template.New("info").Parse(templateText)
	if err != nil {
		return nil, err
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("Could not read build info")
	}

	data := struct {
		Version   string
		GoVersion string
		OSArch    string
		GitCommit string
	}{
		Version:   buildInfo.Main.Version,
		GoVersion: runtime.Version(),
		OSArch:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	out := strings.TrimSpace(buf.String())
	return &out, nil
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Version info",
	Long:    `Detailed information about the version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		msg, err := infoMessage()
		if err != nil {
			log.Error().Msgf("Failed to print message: %s", err)
			return err
		}
		log.Info().Msg(*msg)
		return nil
	},
}
