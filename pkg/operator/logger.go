package operator

import (
	"fmt"

	"inspr.dev/primal/pkg/workflow"
)

type Logger struct {
	*Operator
}

func (op *Operator) NewLogger() *Logger {
	return &Logger{
		op,
	}
}

func (logger *Logger) Task() workflow.Task {
	return workflow.Task{
		ID:    "loggerTask",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
			fmt.Println(logger.Fs)

			self.State = workflow.DONE
		},
	}
}
