package operator

import (
	"fmt"

	"inspr.dev/primal/pkg/filesystem"
	"rogchap.com/v8go"
)

// TODO: modify the structure so VM is initialized and can accept new contexts
// based on files inputed in the application server

// VM defines the vm structure and its fields
type VM struct {
	path string
	file []byte
}

// NewVM returns a reference to a new VM structure with the given path and file
func NewVM(path string, file []byte) *VM {
	return &VM{
		path,
		file,
	}
}

// Handler handles js code in the given filesystem
func (vm *VM) Handler(fs filesystem.FileSystem) {
	iso, _ := v8go.NewIsolate() // creates a new JavaScript VM
	defer iso.Dispose()

	runFn, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		html := fmt.Sprintf("%v", info.Args()[0]) // when the JS function is called this Go callback will execute
		fmt.Println("\n" + html + "\n")
		return nil // you can return a value back to the JS caller if required
	})

	global, _ := v8go.NewObjectTemplate(iso) // a template that represents a JS Object
	global.Set("run", runFn)                 // sets the "print" property of the Object to our function
	ctx1, _ := v8go.NewContext(iso, global)  // new Context with the global Object set to our object template

	_, err := ctx1.RunScript(string(vm.file), "stdin.js")
	if err != nil {
		fmt.Println(err)
	}
}
