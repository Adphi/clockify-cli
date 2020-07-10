package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	configCmd "github.com/pkuebler/clockify-cli/pkg/cmd/config"
	createCmd "github.com/pkuebler/clockify-cli/pkg/cmd/create"
	deleteCmd "github.com/pkuebler/clockify-cli/pkg/cmd/delete"
	listCmd "github.com/pkuebler/clockify-cli/pkg/cmd/list"
	userinfoCmd "github.com/pkuebler/clockify-cli/pkg/cmd/userinfo"
	"github.com/pkuebler/clockify-cli/pkg/config"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

// NewRootCmd will add flags and subcommands to the different commands
func NewRootCmd(goCtx context.Context) *cobra.Command {
	ctx := runtime.NewRuntime(
		goCtx,
		logrus.NewEntry(logrus.StandardLogger()),
	)

	ctx.Log.Logger.SetLevel(logrus.DebugLevel)

	home, err := homedir.Dir()
	if err != nil {
		ctx.Log.Panic(err)
	}
	ctx.ConfigFile = filepath.Join(home, ".clockify-cli.json")

	rootCmd := &cobra.Command{
		Use:  "clockify-cli",
		Long: "cli tool to generate reports from clockify time tracker.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Flags
	rootCmd.PersistentFlags().StringVarP(&ctx.ConfigFile, "config", "c", ctx.ConfigFile, "--config clockify-cli.json")
	rootCmd.PersistentFlags().BoolVarP(&ctx.Interactive, "interactive", "i", false, "--interactive")
	rootCmd.PersistentFlags().StringVarP(&ctx.Output, "output", "o", "", "--output json")
	rootCmd.PersistentFlags().StringVarP(&ctx.WorkspaceID, "workspace", "w", "", "--workspace <id>")

	// Log file
	logPath := filepath.Join(home, ".clockify-log")
	logFile, err := os.Create(logPath)
	if err != nil {
		fmt.Printf("Cloud not open log file: %s\n", logPath)
	}
	ctx.Log.Logger.SetOutput(logFile)

	// Init
	ctx.Config = config.ReadConfig(ctx.ConfigFile, ctx.Log.WithField("component", "config"))
	ctx.Client, err = clockify.NewAPIClient(ctx.Config.APIConfig.APIEndpoint, ctx.Config.APIConfig.ReportEndpoint, ctx.Config.APIConfig.ClockifyKey, nil, ctx.Log.WithField("component", "client"))
	if err != nil {
		fmt.Print(err.Error())
		return rootCmd
	}

	ctx.Client.StartRatelimit(ctx.Context)

	// Commands
	rootCmd.AddCommand(configCmd.NewCmdConfig(ctx))
	rootCmd.AddCommand(userinfoCmd.NewCmdInfo(ctx))
	rootCmd.AddCommand(createCmd.NewCmdCreate(ctx))
	rootCmd.AddCommand(deleteCmd.NewCmdDelete(ctx))
	rootCmd.AddCommand(listCmd.NewCmdList(ctx))

	return rootCmd
}
