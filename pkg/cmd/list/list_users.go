package list

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdListUsers creates a list users command
func NewCmdListUsers(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user",
		Short:   "list users",
		Long:    "list users",
		Example: "user",
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

			opts := &clockify.UserListFilter{
				PageSize: 50,
			}

			entries, err := ctx.Client.User.List(ctx.WorkspaceID, opts)
			if err != nil {
				log.Fatal(err)
			}

			ctx.Log.Tracef("Get %d entries", len(*entries))

			output, _ := json.MarshalIndent(entries, "", "\t")
			fmt.Println(string(output))
		},
	}

	return cmd
}
