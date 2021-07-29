package workflow

import (
	"fmt"
)

// Add adds a new task in a workflow
func (w *Workflow) Add(task Task) {
	w.Tasks = append(w.Tasks, &task)
}

// Run execs a workflow's tasks
func (w *Workflow) Run() {
	for {
		if allTasksAreDone(w.Tasks) {
			fmt.Println("Workflow is done")

			for _, task := range w.Tasks {
				task.State = IDLE
			}
			break
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

func allTasksAreDone(tasks []*Task) bool {
	for _, task := range tasks {
		if task.State != DONE {
			return false
		}
	}

	return true
}
