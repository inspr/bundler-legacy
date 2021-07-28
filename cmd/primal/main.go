package main

import (
	"os"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/platform"
)

func main() {
	path, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()

	// Define Primal with options
	primal := api.Primal{
		Options: api.PrimalOptions{
			Platform: "web",
			Root:     path,
			Watch:    true,
		},
	}

	// Get platform depending on passsed options to Primal
	platform := platform.NewPlatform(primal.Options, fs)
	if primal.Options.Watch {
		platform.Watch()
	} else {
		platform.Run()
	}
}
