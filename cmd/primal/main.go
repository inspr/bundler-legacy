package main

import (
	"context"
	"fmt"
	"os"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/platform/web"
	"inspr.dev/primal/pkg/server"
)

type Compiler struct {
	operators []api.Operator
	spec      api.Spec
}

func (c *Compiler) Add(ops ...api.Operator) *Compiler {
	c.operators = append(c.operators, ops...)
	return c
}

func NewCompiler(root string, fs filesystem.FileSystem) *Compiler {
	cp := &Compiler{
		operators: []api.Operator{},
		spec: api.Spec{
			Root:  root,
			Files: fs,
		},
	}

	return cp
}

func (c *Compiler) Apply() {
	ctx := context.Background()
	spec := c.spec

	// defer cancel()

	for _, op := range c.operators {
		(func() {
			go op.Apply(ctx, spec)

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
	return fmt.Sprint(c.spec.Files)
}

func (c *Compiler) WriteToDisk(dest string) error {
	return c.spec.Files.Flush(dest)
}

func main() {
	root, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()
	primal := NewCompiler(root, fs)

	Bundler := web.NewBundler().WithMinification().Target("client")
	HtmlGen := web.NewHtml()

	primal.
		Add(Bundler, HtmlGen).
		Apply()

	err := primal.WriteToDisk("./__build__")
	fmt.Println(err)
	fmt.Println(primal)

	server.Run(primal.spec.Files)
}
