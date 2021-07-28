package main

import (
	"os"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/platform"
	w "inspr.dev/primal/pkg/workflow"
)

type PrimalOptions struct {
	watch bool
	root  string
}

type Primal struct {
	workflow w.Workflow
	options  PrimalOptions
}

func main() {
	path, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()

	platformOptions := platform.PlatformOptions{
		Platform: "web",
		Root:     path,
	}

	platform := platform.NewPlatform(platformOptions, fs)

	// Define Primal with options
	p := Primal{
		options: PrimalOptions{
			root:  path,
			watch: true,
		},
	}

	// Define Primal main task that start Primal and operators
	primalMain := w.Task{
		ID:    "primal-main-task",
		State: w.IDLE,
		Run: func(t *w.Task) {
			if p.options.watch {
				platform.Watch()
			} else {
				platform.Run()
			}

			t.State = w.DONE
		},
	}

	p.workflow = w.Workflow{
		Tasks: []*w.Task{&primalMain},
	}

	// Start Primal workflow
	p.workflow.Run()
}
