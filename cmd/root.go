package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/config"
)

// rootCmd
var rootCmd = cobra.Command{
	Use:  "clockify-cli",
	Long: "cli tool to generate reports from clockify time tracker.",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, report)
	},
}

// RootCmd will add flags and subcommands to the different commands
func RootCmd() *cobra.Command {
	rootCmd.AddCommand(&reportCmd, &configureCmd)
	return &rootCmd
}

func execWithConfig(cmd *cobra.Command, fn func(config *config.Config, log *logrus.Logger)) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Info("Read Config...")
	config := config.ReadConfig("config.json", log)

	fn(config, log)
}
