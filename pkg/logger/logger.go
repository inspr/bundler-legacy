package logger

import (
	"context"
	"fmt"

	"inspr.dev/primal/pkg/api"
)

type Logger struct {
	meta api.Metadata
}

func NewLogger() *Logger {
	return &Logger{
		meta: api.NewMetadata(),
	}
}

func (h *Logger) Apply(ctx context.Context, opts api.OperatorOptions) error {
	h.meta.State <- api.WORKING
	h.meta.Progress <- 0.0

	fmt.Println(opts.Files)

	h.meta.State <- api.DONE
	h.meta.Progress <- 1.0

	return nil
}

func (h *Logger) Meta() api.Metadata {
	return h.meta
}
