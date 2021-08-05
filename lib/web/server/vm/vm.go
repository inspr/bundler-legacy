package vm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"rogchap.com/v8go"
)

type Request struct {
	UUID uuid.UUID
	Path string
	// TODO: Add the request information such as host, query, path, etc...

	// Add an API to add fetch, store, etc...
}

type Response struct {
	HTML []byte
	// errors []error
}

// VirtualMachine defines the vm structure and its fields
type vm struct {
	ctx context.Context
	iso *v8go.Isolate

	script string
}

// New create a new virtual machine
func New(ctx context.Context) *vm {
	iso, _ := v8go.NewIsolate()

	return &vm{
		ctx: ctx,
		iso: iso,
	}
}

// WithScript set the script to be executed by the virtual machine
func (machine *vm) WithScript(script string) *vm {
	// TODO: Add checks before assing the script
	machine.script = script
	return machine
}

func execVirtualMachine(machine *vm, runCtx Request, repl chan Response) {
	runFn, _ := v8go.NewFunctionTemplate(machine.iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		html := fmt.Sprintf("%v", info.Args()[0]) // when the JS function is called this Go callback will execute
		repl <- Response{
			HTML: []byte(html),
		}

		close(repl)
		return nil // you can return a value back to the JS caller if required
	})

	global, _ := v8go.NewObjectTemplate(machine.iso) // a template that represents a JS Object
	global.Set("run", runFn)                         // sets the "print" property of the Object to our function

	ctx1, _ := v8go.NewContext(machine.iso, global) // new Context with the global Object set to our object template

	_, err := ctx1.RunScript(string(machine.script), runCtx.UUID.String())
	if err != nil {
		fmt.Println(err)
		repl <- Response{}
		close(repl)
	}
}

func (machine *vm) Run(req Request) chan Response {
	repl := make(chan Response)
	go execVirtualMachine(machine, req, repl)
	return repl
}

func (machine *vm) Close() {
	machine.iso.Dispose()
}
