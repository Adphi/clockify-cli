package list

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdListTimeEntries creates a list time-entriess command
func NewCmdListTimeEntries(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "time-entries",
		Short:   "list time-entriess",
		Long:    "list time-entriess",
		Example: "time-entries",
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

			opts := &clockify.TimeEntryListFilter{
				PageSize: 50,
			}

			entries, err := ctx.Client.TimeEntry.List(ctx.WorkspaceID, user.ID, opts)
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
