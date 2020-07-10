package delete

import (
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdDelete creates a delete root command
func NewCmdDelete(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "delete filter",
		Long:    "delete the cli",
		Example: "delete filter",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return cmd
}
