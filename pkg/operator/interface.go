package operator

import (
	"context"

	fs "inspr.dev/primal/pkg/filesystem"
)

type Spec struct {
	// Loaction of the app Root on Primal's FileSysten
	Root string

	// Environment variables
	Enviroment map[string]string
}

// Operator defines a chain of actions to be executed by Primal.
// Think about an operator as a step.
// Primal will look an operator, execute and then call the operators defined by its next func
type Operator interface {
	Apply(ctx context.Context, spec Spec, fs fs.FileSystem) error

	// Progress return a channel of numbers from 0.0 ... 1.0 used to inform the % completed by a task
	Progress() <-chan float32

	// Messages return relevant information about the operator step
	Messages() <-chan string

	// Done inform if an operator fully executed or not
	Done() <-chan bool
}
