package operator

import (
	"context"
)

// Move to interface
type FileSystem interface {

	// Get return the file in a file system, given a path (ext included)
	Get(path string) ([]byte, error)

	// Match return a list of file paths that match the requested regex
	// this function is useful to select files to apply operations
	Match(regex string) ([]string, error)

	// GetAll return the files that match a given regex
	// think of it as a combination of Get and Match
	GetAll(regex string) ([]byte, error)
}

type Spec struct {
	// Loaction of the app Root on Primal's FileSysten
	Root string

	// Environment variables
	Enviroment map[string]string
}

type OperatorProps struct {
	Progress chan float32
	Messages chan string
	Done     chan bool
}

// Task is a function that the operator will execute when called
type Task func(ctx context.Context, spec *Spec, fs *FileSystem, op *OperatorProps)

// Operator defines a chain of actions to be executed by Primal.
// Think about an operator as a step.
// Primal will look an operator, execute and then call the operators defined by its next func
type Operator interface {
	Apply(fn Task) error

	// Progress return a channel of numbers from 0.0 ... 1.0 used to inform the % completed by a task
	Progress() <-chan float32

	// Messages return relevant information about the operator step
	Messages() <-chan string

	// Done inform if an operator fully executed or not
	Done() <-chan bool

	// Next return a list of operators to be applied after the current operator
	Next() []Operator
}
