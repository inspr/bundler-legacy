package platform

import "inspr.dev/primal/pkg/filesystem"

type Platform struct {
	PlatformInterface

	fs      filesystem.FileSystem
	options PlatformOptions
}

type PlatformInterface interface {
	Run()
	Watch()
}

type PlatformOptions struct {
	Platform string
	Root     string
}
