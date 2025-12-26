package cmd

import (
	"os"

	"github.com/dark0dave/wpm/pkg/url"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	urlName, urlPath, urlVersion string
	urlAddCmd                    = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Add url dependencies",
		Long:    `Add url dependencies to a manifest file`,
		Run: func(cmd *cobra.Command, args []string) {
			urlDependency := url.New(urlName, urlPath, urlVersion)
			for _, dep := range m.Dependencies {
				if dep == urlDependency {
					log.Error().Msgf("Url dependency already exists: %+v", dep)
					os.Exit(1)
				}
			}
			viper.Set("dependencies.url", append(m.Dependencies, urlDependency))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
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
			urlDependency := url.New(urlName, urlPath, urlVersion)
			for i, dep := range m.Dependencies {
				if dep == urlDependency {
					viper.Set("dependencies.url", append(m.Dependencies[:i], m.Dependencies[i+1:]...))
					log.Debug().Msgf("Removed url dependency: %+v", dep)
					break
				}
			}
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
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
