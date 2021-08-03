package workflow

import (
	"fmt"
)

var errChan chan error

func init() {
	errChan = make(chan error)
}

// Add adds a new task in a workflow
func (w *Workflow) Add(task Task) {
	task.ErrChan = errChan
	w.Tasks = append(w.Tasks, &task)
}

// Run execs a workflow's tasks
func (w *Workflow) Run() {
	w.ErrChan = errChan
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
					continue Main
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
