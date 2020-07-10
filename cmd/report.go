package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/config"
)

var reportCmd = cobra.Command{
	Use:   "report",
	Short: "Create a workreport",
	Long:  "Create a workreport based on workday settings.",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, report)
	},
}

func report(config *config.Config, log *logrus.Logger) {
	// if default workspace defined?
	// if overwrite by parameter
	// if workspace defined?

	// if worktime defined? by parameter?

	// if filter -> add filter
}
