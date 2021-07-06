package vm

import (
	"fmt"

	"rogchap.com/v8go"
)

func Run(file []byte) {
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

	_, err := ctx1.RunScript(string(file), "stdin.js")
	if err != nil {
		fmt.Println(err)
	}
}
