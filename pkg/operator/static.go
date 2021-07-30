package operator

import (
	"fmt"
	"io/ioutil"
	"path"

	"inspr.dev/primal/pkg/workflow"
)

// Static defines the static file generator structure
type Static struct {
	*Operator
	files []string
}

// NewStatic returns a reference to a new static file generator structure
// with the given path and files
func (operator *Operator) NewStatic(files []string) *Static {
	return &Static{
		Operator: operator,
		files:    files,
	}
}

// Task - Static operator handle function
func (static *Static) Task() workflow.Task {
	return workflow.Task{
		ID:    "static",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
			for _, relativePath := range static.files {
				fullPath := path.Join(static.Options.Root, relativePath)

				data, err := ioutil.ReadFile(fullPath)
				if err != nil {
					fmt.Println(err)
				}

				static.Fs.Write("/"+relativePath, data)
			}

			self.State = workflow.DONE
		},
	}

}
