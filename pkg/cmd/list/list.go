package list

import (
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdList creates a list root command
func NewCmdList(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list filter",
		Long:    "list the cli",
		Example: "list filter",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewCmdListClients(ctx))
	cmd.AddCommand(NewCmdListProjects(ctx))
	cmd.AddCommand(NewCmdListTags(ctx))
	cmd.AddCommand(NewCmdListUsers(ctx))
	cmd.AddCommand(NewCmdListUserGroups(ctx))
	cmd.AddCommand(NewCmdListWorkspaces(ctx))
	cmd.AddCommand(NewCmdListTasks(ctx))
	cmd.AddCommand(NewCmdListTimeEntries(ctx))

	return cmd
}
