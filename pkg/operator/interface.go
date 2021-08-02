package operator

import (
	"context"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

// Operator is created so it's possible to define methods
// for the api.Operator struct in this package
type Operator api.Operator

// MainOperators is a map type of operator ID's to operator structures
type MainOperators map[string]api.OperatorInterface

// MainOps is a structure that contains the operators which are
// used by every platform
var MainOps MainOperators

// NewOperator returns a reference to a new operator with the given
// filesystem and primal options
func NewOperator(options api.PrimalOptions, fs filesystem.FileSystem) *Operator {
	return &Operator{
		Ctx:     context.Background(),
		Fs:      fs,
		Options: options,
	}
}

// InitMainOperators creates the main platform operators and store them on MainOps
func (op *Operator) InitMainOperators() {
	MainOps = map[string]api.OperatorInterface{
		"html": api.OperatorInterface(op.NewHtml()),
		// TODO: create method to add new operators
		// "static": api.OperatorInterface(op.NewStatic([]string{"template/sw.js"})),
		"disk":   api.OperatorInterface(op.NewDisk()),
		"logger": api.OperatorInterface(op.NewLogger()),
	}
}
