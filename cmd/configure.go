package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/config"
)

var configureCmd = cobra.Command{
	Use:   "configure",
	Short: "Configure clockify-cli",
	Long:  "Create a config file and add the api key.",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, configure)
	},
}

func configure(config *config.Config, log *logrus.Logger) {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("API Key must have more than 6 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "API Key",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// add to config file
	fmt.Println(result)

	// crawl workspaces
	// - select default workspace

	// save
}
