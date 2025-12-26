package cmd

import (
	"fmt"
	u "net/url"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			parseUrl, err := u.Parse(urlPath)
			if err != nil {
				return err
			}
			urlDependency := url.New(urlName, urlPath, *parseUrl)
			for _, dep := range m.Dependencies {
				if dep == urlDependency {
					return fmt.Errorf("Url dependency already exists: %+v", dep)
				}
			}
			viper.Set("dependencies.url", append(m.Dependencies, urlDependency))
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
				return err
			}
			log.Debug().Msgf("Written new config: %+v", viper.Get("dependencies.url"))
			return nil
		},
	}
	urlrmCmd = &cobra.Command{
		Use:     "url",
		Aliases: []string{"u"},
		Short:   "Remove url dependencies",
		Long:    `Remove url dependencies to a manifest file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseUrl, err := u.Parse(urlPath)
			if err != nil {
				return err
			}
			urlDependency := url.New(urlName, urlPath, *parseUrl)
			for i, dep := range m.Dependencies {
				if dep == urlDependency {
					viper.Set("dependencies.url", append(m.Dependencies[:i], m.Dependencies[i+1:]...))
					log.Debug().Msgf("Removed url dependency: %+v", dep)
					break
				}
			}
			if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
				log.Error().Msgf("Failed to write to config, %s", err)
				return err
			}
			log.Debug().Msgf("Written new config %+v", viper.Get("dependencies.url"))
			return nil
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
