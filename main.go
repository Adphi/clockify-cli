package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkuebler/clockify-cli/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.NewRootCmd(ctx).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run command: %v\n", err)
		os.Exit(1)
	}
}
