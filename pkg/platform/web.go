package platform

import (
	"os"
	"os/signal"
	"syscall"

	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

// Web defines a web platform data
type Web struct {
	*Platform

	workflow workflow.Workflow
}

// Web returns a web platform with it's tasks
func (p *Platform) Web() PlatformInterface {
	web := &Web{
		Platform: p,
	}

	for _, ops := range operator.MainOps {
		web.workflow.Add(ops.Task())
	}

	return web
}

// Run executes the workflow for the web platform
func (w *Web) Run() {
	w.Bundler.Target("client").Build()
	w.workflow.Run()
}

// Watch executes the workflow for the web platform in watch mode
func (w *Web) Watch() {
	w.Bundler.Target("client").Watch()
	w.workflow.Run()

	go Start(w.Fs)
	GracefullShutdown()
}

// ! Should this be here?
// GracefullShutdown
func GracefullShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	os.Exit(1)
}
