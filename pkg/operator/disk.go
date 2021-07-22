package operator

import (
	"os"

	"inspr.dev/primal/pkg/workflow"
)

type Disk struct {
	*Operator
}

func (op *Operator) NewDisk() *Disk {
	return &Disk{
		op,
	}
}

func (d *Disk) Task() workflow.Task {
	return workflow.Task{
		ID:    "diskTask",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
			path := d.root + "/__build__"
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
				os.Mkdir(path+"/assets", 0755)
			}

			for key, file := range d.fs.Raw() {
				// TODO: catch the error and return in an "errors" chan
				f, _ := os.Create(path + key)
				f.Write(file)
			}

			self.State = workflow.DONE
		},
	}
}
