package main

import (
	"context"
	"fmt"
	"os"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/disk"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/logger"
	"inspr.dev/primal/pkg/platform/web"
	"inspr.dev/primal/pkg/server"
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
		}

		runOperator := func() {
			go op.Apply(ctx, opts)

			for {
				select {
				case state := <-op.Meta().State:
					if state == api.DONE {
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

// func (c *Compiler) WriteToDisk(dest string) error {
// 	return c.Files.Flush(dest)
// }

func main() {
	root, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()
	primal := NewCompiler(root, fs)

	Bundler := web.NewBundler().WithMinification().Target("client")
	HtmlGen := web.NewHtml()
	Disk := disk.NewDisk("./__build2__")
	Server := server.NewServer(3049)
	Logger := logger.NewLogger()

	primal.
		Add(Bundler).
		Add(HtmlGen).
		Add(Logger, Disk, Server).
		Apply()
}
