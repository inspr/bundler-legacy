package workflow

import (
	"fmt"
)

const (
	IDLE Status = 1 << iota
	RUNNING
	DONE
)

type Status int

type Task struct {
	ID     string
	Before func()
	After  func()

	Run func(*Task)

	DependsOn []*Task

	State Status
}

// Workflow is a set of tasks with a predefined order of execution
type WorkflowInterface interface {
	Add(Task)
	Remove(Task)

	Run()
}

type Workflow struct {
	// WorkflowInterface

	Tasks []*Task
}

func (w *Workflow) Add(task Task) {
	w.Tasks = append(w.Tasks, &task)
}

func allTasksAreDone(tasks []*Task) bool {
	for _, task := range tasks {
		if task.State != DONE {
			return false
		}
	}

	return true
}

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
