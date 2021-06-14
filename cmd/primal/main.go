package main

import (
	"context"
	"fmt"
	"os"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/platform/web"
)

type Compiler struct {
	fs        filesystem.FileSystem
	operators []operator.Operator
	spec      operator.Spec
}

func (c *Compiler) Add(op operator.Operator) {
	c.operators = append(c.operators, op)
}

func NewCompiler(root string, fs filesystem.FileSystem) *Compiler {
	cp := &Compiler{
		fs:        fs,
		operators: []operator.Operator{},
		spec: operator.Spec{
			Root: root,
		},
	}

	return cp
}

func (c *Compiler) Apply() {
	for _, op := range c.operators {
		(func() {
			go op.Apply(context.Background(), c.spec, c.fs)

			for {
				select {
				case <-op.Done():
					return
				case msg := <-op.Messages():
					fmt.Println(msg)
				case v := <-op.Progress():
					fmt.Println(v)
				}
			}
		})()
	}
}

func (c *Compiler) String() string {
	return fmt.Sprint(c.fs)
}

func main() {
	root, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()
	cpl := NewCompiler(root, fs)

	cpl.Add(web.NewBundler().WithDevelopMode().Target("client"))
	cpl.Add(web.NewHtml())

	cpl.Apply()

	fmt.Println(cpl)

	err := cpl.fs.Flush("./__build__")
	fmt.Println(err)
}
