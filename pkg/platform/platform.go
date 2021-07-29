package platform

import (
	"fmt"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/bundler"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
)

// Platform is created so it's possible to define methods
// for the api.Platform struct in this package
type Platform api.Platform

// NewPlatform returns a PlatformInterface of the given platform type
func NewPlatform(options api.PrimalOptions, fs filesystem.FileSystem) (api.PlatformInterface, error) {
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
