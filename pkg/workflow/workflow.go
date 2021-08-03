package workflow

import (
	"fmt"
)

// NewWorkflow return new Workflow struct
func NewWorkflow() *Workflow {
	return &Workflow{
		ErrChan: make(chan error),
	}
}

// Add adds a new task in a workflow
func (w *Workflow) Add(task Task) {
	task.ErrChan = w.ErrChan
	w.Tasks = append(w.Tasks, &task)
}

// Run execs a workflow's tasks
func (w *Workflow) Run() {
Main:
	for {
		select {
		case err := <-w.ErrChan:
			fmt.Println(err)
			break Main
		default:
			if allTasksAreDone(w.Tasks) {
				fmt.Println("Workflow is done")

				for _, task := range w.Tasks {
					task.State = IDLE
				}
				break Main
			}

			for _, task := range w.Tasks {
				if task.State != IDLE {
					continue
				}

				// check if parents are done
				if allTasksAreDone(task.DependsOn) {
					task.State = RUNNING
					go task.Run(task)
				}
			}
		}
	}
}

func allTasksAreDone(tasks []*Task) bool {
	for _, task := range tasks {
		if task.State != DONE {
			return false
		}
	}

	return true
}
