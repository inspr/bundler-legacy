package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID     string
	Before func()
	After  func()

	Run func(*Task)

	DependsOn []*Task

	State string
}

// Workflow is a set of tasks with a predefined order of execution
type Workflow interface {
	Add(Task)
	Remove(Task)

	Run()
}

type workflow struct {
	tasks []*Task
}

func allTasksAreDone(tasks []*Task) bool {
	for _, task := range tasks {
		if task.State != "done" {
			return false
		}
	}

	return true
}

func (w *workflow) Run() {
	for {
		if allTasksAreDone(w.tasks) {
			fmt.Println("Workflow is done")
			break
		}

		for _, task := range w.tasks {
			if task.State != "idle" {
				continue
			}

			// check if parents are done
			if allTasksAreDone(task.DependsOn) {
				task.State = "running"
				go task.Run(task)
			}
		}
	}
}

func main() {
	w := workflow{}

	t1 := Task{
		ID:    "t1",
		State: "idle",
		Run: func(self *Task) {
			<-time.After(1 * time.Second)
			fmt.Println("Task 1 Ready")
			self.State = "done"
		},
	}

	t4 := Task{
		ID:    "t4",
		State: "idle",
		Run: func(self *Task) {
			<-time.After(20 * time.Second)
			fmt.Println("Task 4 Ready")
			self.State = "done"
		},
	}

	t2 := Task{
		ID:    "t2",
		State: "idle",
		Run: func(self *Task) {
			fmt.Println("Task 2 Ready")
			self.State = "done"
		},
		DependsOn: []*Task{&t1},
	}

	t3 := Task{
		ID:    "t3",
		State: "idle",
		Run: func(self *Task) {
			<-time.After(2 * time.Second)
			fmt.Println("Task 3 Ready")
			self.State = "done"
		},
		DependsOn: []*Task{&t1, &t2},
	}

	t6 := Task{
		ID:    "t6",
		State: "idle",
		Run: func(self *Task) {
			<-time.After(4 * time.Second)
			fmt.Println("Task 6 Ready")
			self.State = "done"
		},
	}

	w2 := workflow{
		tasks: []*Task{&t6},
	}

	t7 := Task{
		ID:    "sub-workflow",
		State: "idle",
		Run: func(t *Task) {
			w2.Run()
			fmt.Println("Subworkflow 2 Ready")
			t.State = "done"
		},
	}

	w.tasks = append(w.tasks, []*Task{&t2, &t1, &t4, &t3, &t7}...)
	w.Run()
}
