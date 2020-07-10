package create

import (
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdCreate creates a create root command
func NewCmdCreate(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "create filter",
		Long:    "create the cli",
		Example: "create filter",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return cmd
}
