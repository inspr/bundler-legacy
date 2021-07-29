package workflow

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
