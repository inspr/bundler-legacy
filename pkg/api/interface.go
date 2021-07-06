package api

import (
	"context"

	fs "inspr.dev/primal/pkg/filesystem"
)

type OperatorProps struct {
	Context context.Context
	Files   fs.FileSystem
}

type OperatorOptions struct {
	Root string

	Watch bool

	// Environment variables
	Enviroment map[string]string
}

// Operator defines a chain of actions to be executed by Primal.
// Think about an operator as a step.
// Primal will look an operator, execute and then call the operators defined by its next func
type Operator interface {
	Apply(props OperatorProps, opts OperatorOptions)
	Meta() Metadata
}
