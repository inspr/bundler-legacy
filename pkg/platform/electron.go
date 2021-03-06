package platform

import (
	"context"
	"fmt"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/operator"
)

// Electron defines an electron platform data
type Electron struct {
	*Platform
}

// Electron returns an electron platform with it's tasks
func (p *Platform) Electron() api.PlatformInterface {
	electron := &Electron{
		Platform: p,
	}

	for _, ops := range operator.MainOps {
		electron.Platform.Workflow.Add(ops.Task())
	}

	return electron
}

// Run executes the workflow for the electron platform
func (e *Electron) Run() {
	fmt.Println("Implement me.")
}

// Watch executes the workflow for the electron platform in watch mode
func (e *Electron) Watch(ctx context.Context, cancel context.CancelFunc) {
	fmt.Println("Implement me.")
}
