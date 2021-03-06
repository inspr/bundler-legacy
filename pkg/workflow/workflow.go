package workflow

import (
	"context"
	"fmt"
)

// NewWorkflow return new Workflow struct
func NewWorkflow() *Workflow {
	return &Workflow{
		Tasks:   map[string]*Task{},
		ErrChan: make(chan error),
	}
}

// Add adds a new task in a workflow
func (w *Workflow) Add(task Task) {
	task.ErrChan = w.ErrChan
	w.Tasks[task.ID] = &task
}

// Run execs a workflow's tasks
func (w *Workflow) Run(ctx context.Context, cancel context.CancelFunc) {
Main:
	for {
		select {
		case err := <-w.ErrChan:
			fmt.Println(err)
			cancel()
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
				if allTasksAreDone(task.Dependencies) {
					task.State = RUNNING
					go task.Run(ctx, task)
				}
			}
		}
	}
}

func allTasksAreDone(tasks map[string]*Task) bool {
	for _, task := range tasks {
		if task.State != DONE {
			return false
		}
	}

	return true
}

func (t *Task) DependsOn(tasks ...*Task) {
	if len(t.Dependencies) == 0 {
		t.Dependencies = map[string]*Task{}
	}
	for _, task := range tasks {
		t.Dependencies[task.ID] = task
	}
}
