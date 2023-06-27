// Package config implements all commands of KoboMail
package commands

import (
	"fmt"
	"os"

	"github.com/bjw-s/lego-auto/internal/config"
	"github.com/spf13/cobra"
)

var (
	conf = &config.Config{}

	rootCmd = &cobra.Command{
		Use:   "kobomail",
		Short: "lego-auto automates the process of generating LetsEncrypt certificates",
		Long: `lego-auto automates the process of generating LetsEncrypt certificates.
More information available at the Github Repo (https://github.com/bjw-s/lego-auto)`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringP("config", "c", "", "config.yaml file")
}

func initConfig() {
	var err error
	conf, err = config.LoadConfig(rootCmd.PersistentFlags())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := conf.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
