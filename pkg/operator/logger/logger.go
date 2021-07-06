package logger

import (
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

func (h *Logger) Apply(props api.OperatorProps, opts api.OperatorOptions) {

	var log = func() {
		fmt.Println(props.Files)
		h.meta.Done <- true
	}

	log()
Main:
	for {
		select {
		case <-h.meta.Close:
			break Main

		case <-h.meta.Refresh:
			log()
		}
	}
}

func (h *Logger) Meta() api.Metadata {
	return h.meta
}
