package list

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdListTasks creates a list tasks command
func NewCmdListTasks(ctx *runtime.Runtime) *cobra.Command {
	projectID := ""

	cmd := &cobra.Command{
		Use:     "task",
		Short:   "list tasks",
		Long:    "list tasks",
		Example: "task",
		Run: func(cmd *cobra.Command, args []string) {
			// request user data
			user, err := ctx.Client.User.Info()
			if err != nil {
				log.Fatal(err)
				return
			}

			// workspace
			if ctx.WorkspaceID == "" {
				ctx.WorkspaceID = user.DefaultWorkspace
			}

			if projectID == "" {
				ctx.Log.Error("Need ProjectID.")
				fmt.Println("error: need --projectID argument.")
				return
			}

			opts := &clockify.TaskListFilter{
				PageSize: 50,
			}

			entries, err := ctx.Client.Task.List(ctx.WorkspaceID, projectID, opts)
			if err != nil {
				log.Fatal(err)
			}

			ctx.Log.Tracef("Get %d entries", len(*entries))

			output, _ := json.MarshalIndent(entries, "", "\t")
			fmt.Println(string(output))
		},
	}

	// flags
	cmd.Flags().StringVarP(&projectID, "projectID", "p", "", "--projectID 5b1e6b160cb8793dd93ec120")

	return cmd
}
