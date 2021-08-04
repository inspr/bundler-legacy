package workflow

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkflow_Add(t *testing.T) {
	type args struct {
		task Task
	}
	tests := []struct {
		name string
		w    *Workflow
		args args
	}{
		{
			name: "append task in workflow",
			w:    generateMockWorkflow(),
			args: args{
				Task{
					ID: "newTask",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numTasks := len(tt.w.Tasks)
			tt.w.Add(tt.args.task)
			if len(tt.w.Tasks) != numTasks+1 {
				t.Errorf("task wasn't added")
			}
		})
	}
}

func TestWorkflow_Run(t *testing.T) {
	tests := []struct {
		name  string
		w     *Workflow
		check func([]*Task) error
	}{
		{
			name: "run a workflow with all tasks done",
			w:    generateMockWorkflow(),
			check: func(t []*Task) error {
				for _, task := range t {
					if task.State != IDLE {
						return fmt.Errorf("all tasks should be IDLE")
					}
				}
				return nil
			},
		},
		{
			name: "run a workflow with IDLE task with dependencies",
			w:    mockWorkflowWithDependentTasks(),
			check: func(t []*Task) error {
				for _, task := range t {
					if task.State == IDLE {
						return fmt.Errorf("tasks shouldn't be IDLE")
					}
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go tt.w.Run()
			time.Sleep(time.Microsecond * 500)
			err := tt.check(tt.w.Tasks)
			if err != nil {
				t.Errorf("error in Run(): %v\n", err)
			}
		})
	}
}

// Auxiliar methods

func generateMockWorkflow() *Workflow {
	w := &Workflow{}

	w.Add(Task{
		ID:    "task1",
		State: DONE,
	})
	w.Add(Task{
		ID:    "task2",
		State: DONE,
	})

	return w
}

func mockWorkflowWithDependentTasks() *Workflow {
	wf := Workflow{}

	wf.Add(Task{
		ID:    "task2",
		State: RUNNING,
	})
	wf.Add(Task{
		ID:    "task3",
		State: IDLE,
		DependsOn: []*Task{
			{
				ID:    "task1",
				State: DONE,
			},
		},
		Run: func(self *Task) {
			self.State = DONE
		},
	})

	return &wf
}
