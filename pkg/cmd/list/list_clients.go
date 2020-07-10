package list

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdListClients creates a list clients command
func NewCmdListClients(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "client",
		Short:   "list clients",
		Long:    "list clients",
		Example: "client",
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

			opts := &clockify.ClientListFilter{
				PageSize: 50,
			}

			entries, err := ctx.Client.Client.List(ctx.WorkspaceID, opts)
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
