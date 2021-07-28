package platform

import (
	"fmt"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/bundler"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
)

type Platform api.Platform

type PlatformInterface api.PlatformInterface

func NewPlatform(options api.PrimalOptions, fs filesystem.FileSystem) PlatformInterface {
	operator.NewOperator(options, fs).InitMainOperators()

	platform := &Platform{
		Bundler: bundler.NewBundler(options.Root, fs),
		Options: options,
		Fs:      fs,
	}

	switch options.Platform {
	case "web":
		return platform.Web()
	case "electron":
		return platform.Electron()
	default:
		err := fmt.Sprintf("platform '%s' is not supported.", options.Platform)
		panic(err)
	}
}
