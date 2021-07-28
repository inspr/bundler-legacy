package operator

import (
	"context"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

type Operator api.Operator

type OperatorInterface api.OperatorInterface

type MainOperators map[string]OperatorInterface

var MainOps MainOperators

func NewOperator(options api.PrimalOptions, fs filesystem.FileSystem) *Operator {
	return &Operator{
		Ctx:     context.Background(),
		Fs:      fs,
		Options: options,
	}
}

func (op *Operator) InitMainOperators() {
	MainOps = map[string]OperatorInterface{
		"html":   OperatorInterface(op.NewHtml()),
		"disk":   OperatorInterface(op.NewDisk()),
		"logger": OperatorInterface(op.NewLogger()),
	}
}
