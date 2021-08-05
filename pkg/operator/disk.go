package operator

import (
	"os"

	"inspr.dev/primal/pkg/workflow"
)

// Disk is the logger operator
type Disk struct {
	*Operator
}

// NewDisk returns a new disk operator
func (op *Operator) NewDisk() *Disk {
	return &Disk{
		op,
	}
}

// Task returns a disk operator's workflow task
func (disk *Disk) Task() workflow.Task {
	return workflow.Task{
		ID:    "disk",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
			path := disk.Options.Root + "/__build__"
			dir := disk.Options.Root + "/template"
			if _, err := os.Stat(dir); err == nil {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 0755)
					os.Mkdir(path+"/assets", 0755)
				}

				for key, file := range disk.Fs.Raw() {
					f, err := os.Create(path + key)
					if err != nil {
						self.ErrChan <- err
					}

					f.Write(file)
				}

				self.State = workflow.DONE
			} else if os.IsNotExist(err) {
				self.ErrChan <- err
			} else {
				panic(err)
			}
		},
	}
}
