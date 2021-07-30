package operator

import (
	"context"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

func NewMockOperator(options api.PrimalOptions) *Operator {
	return &Operator{
		Ctx:     context.Background(),
		Fs:      filesystem.NewMemoryFs(),
		Options: options,
	}
}
