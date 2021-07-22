package operator

import (
	"context"

	"inspr.dev/primal/pkg/filesystem"
)

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
