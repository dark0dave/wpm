package cmd

import (
	"os"

	"github.com/dark0dave/wpm/pkg/manifest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	urlName    string
	urlPath    string
	urlVersion string
	urlAddCmd  = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Add url dependencies",
		Long:    `Add url dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			urlDependency := manifest.UrlDependancy{
				Name:    urlName,
				Path:    urlPath,
				Version: urlVersion,
			}
			for _, dep := range m.Dependencies.UrlDependencies {
				if dep == urlDependency {
					log.Error().Msgf("Url dependency already exists: %+v", dep)
					os.Exit(1)
				}
			}
			viper.Set("dependencies.url", append(m.Dependencies.UrlDependencies, urlDependency))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().AnErr("Failed to write to config", err)
			}
			log.Debug().Msgf("Written new config: %+v", viper.Get("dependencies.url"))
		},
	}
	urlrmCmd = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Remove url dependencies",
		Long:    `Remove url dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			urlDependency := manifest.UrlDependancy{
				Name:    urlName,
				Path:    urlPath,
				Version: urlVersion,
			}
			for i, dep := range m.Dependencies.UrlDependencies {
				if dep == urlDependency {
					viper.Set("dependencies.url", append(m.Dependencies.UrlDependencies[:i], m.Dependencies.UrlDependencies[i+1:]...))
					log.Debug().Msgf("Removed url dependency: %+v", dep)
					break
				}
			}
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().AnErr("Failed to write to config", err)
			}
			log.Debug().Msgf("Written new config %+v", viper.Get("dependencies.url"))
		},
	}
)

func init() {
	urlAddCmd.PersistentFlags().StringVarP(&urlName, "name", "n", "", "")
	urlAddCmd.PersistentFlags().StringVarP(&urlPath, "path", "p", "", "")
	urlAddCmd.PersistentFlags().StringVarP(&urlVersion, "version", "v", "", "")
	urlrmCmd.PersistentFlags().StringVarP(&urlName, "name", "n", "", "")
	urlrmCmd.PersistentFlags().StringVarP(&urlPath, "path", "p", "", "")
	urlrmCmd.PersistentFlags().StringVarP(&urlVersion, "version", "v", "", "")
	addCmd.AddCommand(urlAddCmd)
	rmCmd.AddCommand(urlrmCmd)
}
