package workflow

import (
	"fmt"
)

// Status is an enum for possible Operators run status
type Status int

const (
	IDLE Status = 1 << iota
	RUNNING
	DONE
)

// WorkflowInterface defines methods to manage and run a workflow
type WorkflowInterface interface {
	Add(Task)
	Remove(Task)
	Run()
}

// Task defines an operator's execution methods and its needed data
type Task struct {
	ID     string
	Before func()
	After  func()

	Run func(*Task)

	DependsOn []*Task

	State Status
}

// Workflow is a set of tasks with a predefined order of execution
type Workflow struct {
	Tasks []*Task
}

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
