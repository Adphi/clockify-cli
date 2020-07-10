package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Netflix/go-env"
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/config"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewCmdConfigAPI creates a config root command
func NewCmdConfigAPI(ctx *runtime.Runtime) *cobra.Command {
	tmp := config.APIConfig{
		APIEndpoint:    "https://api.clockify.me/api/v1",
		ReportEndpoint: "https://reports.api.clockify.me/v1",
	}

	cmd := &cobra.Command{
		Use:     "api",
		Short:   "set api access key and endpoint",
		Long:    "set api access key and endpoint",
		Example: "config api",
		Run: func(cmd *cobra.Command, args []string) {
			// set by env variables, overwrite flags
			env.UnmarshalFromEnviron(tmp)

			// todo: test

			ctx.Config.APIConfig = tmp

			// save
			file, _ := json.MarshalIndent(ctx.Config, "", "\t")
			if err := ioutil.WriteFile(ctx.ConfigFile, file, 0644); err != nil {
				ctx.Log.Error(err)
			}
		},
	}

	cmd.Flags().StringVarP(&tmp.APIEndpoint, "endpoint", "e", ctx.Config.APIConfig.APIEndpoint, "--endpoint https://api.clockify.me/api/v1")
	cmd.Flags().StringVarP(&tmp.ReportEndpoint, "reportEndpoint", "r", ctx.Config.APIConfig.ReportEndpoint, "--endpoint https://reports.api.clockify.me/v1")
	cmd.Flags().StringVarP(&tmp.ClockifyKey, "key", "k", "", "--key XXXXX")

	return cmd
}
