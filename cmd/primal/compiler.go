package main

import (
	"context"
	"fmt"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

type Compiler struct {
	operators []api.Operator
	Root      string
	Files     filesystem.FileSystem
}

func (c *Compiler) Add(ops ...api.Operator) *Compiler {
	c.operators = append(c.operators, ops...)
	return c
}

func NewCompiler(root string, fs filesystem.FileSystem) *Compiler {
	cp := &Compiler{
		operators: []api.Operator{},
		Root:      root,
		Files:     fs,
	}

	return cp
}

func (c *Compiler) Apply() {
	ctx := context.Background()

	for _, op := range c.operators {
		opts := api.OperatorOptions{
			Root:       c.Root,
			Enviroment: make(map[string]string),
			Files:      c.Files,
			Watch:      true,
		}

		runOperator := func() {
			go op.Apply(ctx, opts)

			for {
				select {
				case state := <-op.Meta().State:
					if state == api.DONE {
						fmt.Println("DONE")
						return
					}

					if state == api.READY {
						continue
					}

					if state == api.WORKING {
						fmt.Println("Working")
					}

				case msg := <-op.Meta().Messages:
					fmt.Println(msg)
				case v := <-op.Meta().Progress:
					fmt.Println(v)
				}
			}
		}

		runOperator()
	}
}

func (c *Compiler) String() string {
	return fmt.Sprint(c.Files)
}
