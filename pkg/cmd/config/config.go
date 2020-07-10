package config

import (
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdConfig creates a config root command
func NewCmdConfig(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "config",
		Short:   "config api|report",
		Long:    "Configure the cli",
		Example: "config api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewCmdConfigAPI(ctx))

	return cmd
}
