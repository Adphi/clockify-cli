package runtime

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/config"
)

// Runtime context
type Runtime struct {
	Log     *logrus.Entry
	Config  *config.Config
	Client  *clockify.APIClient
	Context context.Context

	ConfigFile  string
	Interactive bool
	Output      string
	WorkspaceID string
}

// NewRuntime returns a runtime
func NewRuntime(ctx context.Context, log *logrus.Entry) *Runtime {
	return &Runtime{
		Log:     log,
		Context: ctx,

		Interactive: false,
		Output:      "",
	}
}
