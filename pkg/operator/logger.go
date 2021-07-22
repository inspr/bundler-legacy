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

func (l *Logger) Task() workflow.Task {
	return workflow.Task{
		ID:    "loggerTask",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
			fmt.Println(l.fs)

			self.State = workflow.DONE
		},
	}
}
