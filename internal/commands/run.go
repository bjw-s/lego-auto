// Package config implements all commands of lego-auto
package commands

import (
	"github.com/bjw-s/lego-auto/internal/lego_auto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run KoboMail processing",
	Long:  "Run KoboMail processing.",
	RunE: func(cmd *cobra.Command, args []string) error {
		lego_auto.AppConfig = conf
		lego_auto.Run()
		return nil
	},
}
