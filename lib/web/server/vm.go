package operator

import (
	"context"
	"net/http"

	"rogchap.com/v8go"
)

// VM defines the vm structure and its fields
type VM struct {
	mux *http.ServeMux
	ctx context.Context
	iso *v8go.Isolate
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// NewVM returns a reference to a new VM structure with the given path and file
func NewVM(ctx context.Context) *VM {
	iso, _ := v8go.NewIsolate()
	mux := http.NewServeMux()
	mux.HandleFunc("/vm", VMHandler())

	return &VM{
		mux: mux,
		ctx: ctx,
		iso: iso,
	}
}

// Handler handles js code in the given filesystem
func VMHandler() HandlerFunc {

	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	// iso, _ := v8go.NewIsolate() // creates a new JavaScript VM
	// defer iso.Dispose()

	// runFn, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
	// 	html := fmt.Sprintf("%v", info.Args()[0]) // when the JS function is called this Go callback will execute
	// 	// ! this generates the html file to be served by the server
	// 	fmt.Println("\n" + html + "\n")
	// 	return nil // you can return a value back to the JS caller if required
	// })

	// global, _ := v8go.NewObjectTemplate(iso) // a template that represents a JS Object
	// // ! every html page must have a "run" declared in the js file
	// global.Set("run", runFn)                // sets the "print" property of the Object to our function
	// ctx1, _ := v8go.NewContext(iso, global) // new Context with the global Object set to our object template

	// _, err := ctx1.RunScript(string(vm.file), "stdin.js")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func (vm *VM) Run() {

}
