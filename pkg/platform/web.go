package platform

import (
	"os"
	"os/signal"
	"syscall"

	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

type Web struct {
	*Platform

	workflow workflow.Workflow
}

func (p *Platform) Web() PlatformInterface {
	web := &Web{
		Platform: p,
	}

	for _, ops := range operator.MainOps {
		web.workflow.Add(ops.Task())
	}

	return web
}

func (w *Web) Run() {
	w.Bundler.Target("client").Build()
	w.workflow.Run()
}

func (w *Web) Watch() {
	w.Bundler.Target("client").Watch()
	w.workflow.Run()

	go Start(w.Fs)
	GracefullShutdown()
}

func GracefullShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	os.Exit(1)
}
