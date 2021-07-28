package operator

import (
	"context"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/workflow"
)

type OperatorInterface interface {
	Task() workflow.Task
}

type Operator struct {
	ctx  context.Context
	fs   filesystem.FileSystem
	root string
}

func NewOperator(fs filesystem.FileSystem, root string) *Operator {
	return &Operator{
		ctx:  context.Background(),
		fs:   fs,
		root: root,
	}
}

type MainOperators map[string]OperatorInterface

var MainOps MainOperators

func (op *Operator) InitMainOperators() {
	MainOps = map[string]OperatorInterface{
		"html":   OperatorInterface(op.NewHtml()),
		"disk":   OperatorInterface(op.NewDisk()),
		"logger": OperatorInterface(op.NewLogger()),
	}
}
