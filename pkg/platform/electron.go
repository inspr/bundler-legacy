package platform

import (
	"fmt"

	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

type Electron struct {
	*Platform

	workflow workflow.Workflow
}

func (p *Platform) Electron() PlatformInterface {
	electron := &Electron{
		Platform: p,
	}

	for _, ops := range operator.MainOps {
		electron.workflow.Add(ops.Task())
	}

	return electron
}

func (e *Electron) Run() {
	fmt.Println("Implement me.")
}
func (e *Electron) Watch() {
	fmt.Println("Implement me.")
}
