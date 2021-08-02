package platform

import (
	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

// Web defines a web platform data
type Web struct {
	*Platform

	workflow workflow.Workflow
}

// Web returns a web platform with it's tasks
func (p *Platform) Web() api.PlatformInterface {
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

	server := Server{
		reload: w.Bundler.Refresh(),
	}

	go server.Start(w.Fs)
	GracefullShutdown()
}
