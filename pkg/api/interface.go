package api

import (
	"context"

	"inspr.dev/primal/pkg/bundler"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/workflow"
)

// PlatformInterface defines methods that exec a platform
type PlatformInterface interface {
	Run()
	Watch()
}

// OperatorInterface defines methods that returns an operator operation
type OperatorInterface interface {
	Task() workflow.Task
}

// PrimalOptions contains the main information needed for Primal to run
type PrimalOptions struct {
	Root     string
	Platform string
	Watch    bool
}

// Primal contains PrimalOptions data
type Primal struct {
	Workflow workflow.Workflow // TODO: not used now
	Options  PrimalOptions
}

// Operator contains the necessary information for an operator to run
type Operator struct {
	Options PrimalOptions

	Ctx context.Context
	Fs  filesystem.FileSystem
}

// Platform implements how a platform execs and has the needed data for it to run
type Platform struct {
	PlatformInterface

	Options PrimalOptions
	Fs      filesystem.FileSystem
	Bundler *bundler.Bundler
}
