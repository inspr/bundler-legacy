package operator

import (
	"os"

	"inspr.dev/primal/pkg/filesystem"
)

type Disk struct {
	path string
}

func NewDisk(path string) *Disk {
	return &Disk{
		path: path,
	}
}

func (d *Disk) Handler(fs filesystem.FileSystem) {
	path := d.path + "/__build__"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		os.Mkdir(path+"/assets", 0755)
	}

	for key, file := range fs.Raw() {
		// TODO: catch the error and return in an "errors" chan
		f, _ := os.Create(path + key)
		f.Write(file)
	}
}
