package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdInfo creates a userinfo command
func NewCmdInfo(ctx *runtime.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "userinfo",
		Short:   "userinfo blabla",
		Long:    "blablabla",
		Example: "userinfo",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := ctx.Client.User.Info()
			if err != nil {
				ctx.Log.Error(err)
			}

			output, _ := json.MarshalIndent(info, "", "\t")
			fmt.Println(string(output))
		},
	}

	return cmd
}
