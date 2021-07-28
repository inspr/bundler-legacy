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

func NewPlatform(options api.PrimalOptions, fs filesystem.FileSystem) (PlatformInterface, error) {
	operator.NewOperator(options, fs).InitMainOperators()

	platform := &Platform{
		Bundler: bundler.NewBundler(options.Root, fs),
		Options: options,
		Fs:      fs,
	}

	switch options.Platform {
	case "web":
		return platform.Web(), nil
	case "electron":
		return platform.Electron(), nil
	default:
		err := fmt.Errorf("platform '%s' is not supported", options.Platform)
		return nil, err
	}
}
