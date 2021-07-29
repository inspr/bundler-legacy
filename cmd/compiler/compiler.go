package main

import (
	"context"
	"fmt"
	"sync"

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
	opts := api.OperatorOptions{
		Root:       c.Root,
		Enviroment: make(map[string]string),
		Watch:      true,
	}

	var wg sync.WaitGroup

	// Main:
	for idx, op := range c.operators {
		wg.Add(1)

		props := api.OperatorProps{
			Context: context.Background(),
			Files:   c.Files,
		}

		go op.Apply(props, opts)

		go func(idx int, op api.Operator) {
			for {
				select {
				case <-op.Meta().Updated:
					fmt.Println("updated")
					for idx2, opx := range c.operators {
						if idx != idx2 {
							opx.Meta().Refresh <- true
						}
					}

				case msg := <-op.Meta().Messages:
					fmt.Println(msg)

				case <-op.Meta().Done:
					op.Meta().Close <- true
					wg.Done()
					return
				}
			}
		}(idx, op)
	}

	wg.Wait()
}

func (c *Compiler) String() string {
	return fmt.Sprint(c.Files)
}
