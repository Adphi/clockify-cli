package list

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdListWorkspaces creates a list workspacess command
func NewCmdListWorkspaces(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workspace",
		Short:   "list workspaces",
		Long:    "list workspaces",
		Example: "workspace",
		Run: func(cmd *cobra.Command, args []string) {
			entries, err := ctx.Client.Workspace.List()
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
