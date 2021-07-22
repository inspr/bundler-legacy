package operator

import (
	"os"

	"inspr.dev/primal/pkg/filesystem"
)

type VM struct {
	path string
	file []byte
}

func NewVM(path string, file []byte) *VM {
	return &VM{
		path,
		file,
	}
}

func (vm *VM) Handler(fs filesystem.FileSystem) {
	path := vm.path + "/__build__"
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
