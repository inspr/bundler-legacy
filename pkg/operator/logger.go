package operator

import (
	"context"
	"fmt"

	"inspr.dev/primal/pkg/workflow"
)

// Logger is the logger operator
type Logger struct {
	*Operator
}

// NewLogger returns a new logger operator
func (op *Operator) NewLogger() *Logger {
	return &Logger{
		op,
	}
}

// Task returns a logger operator's workflow task
func (logger *Logger) Task() workflow.Task {
	return workflow.Task{
		ID:    "logger",
		State: workflow.IDLE,
		Run: func(ctx context.Context, self *workflow.Task) {
			fmt.Println(logger.Fs)

			self.State = workflow.DONE
		},
	}
}
