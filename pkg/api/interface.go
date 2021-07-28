package api

import (
	"context"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/workflow"
)

type PrimalOptions struct {
	Root     string
	Platform string
	Watch    bool
}

type Primal struct {
	Workflow workflow.Workflow // TODO: not used now
	Options  PrimalOptions
}

type Operator struct {
	Options PrimalOptions

	Ctx context.Context
	Fs  filesystem.FileSystem
}

type OperatorInterface interface {
	Task() workflow.Task
}

type Platform struct {
	PlatformInterface

	Options PrimalOptions
	Fs      filesystem.FileSystem
}

type PlatformInterface interface {
	Run()
	Watch()
}
