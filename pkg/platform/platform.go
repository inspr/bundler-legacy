package platform

import (
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
)

func NewPlatform(options PlatformOptions, fs filesystem.FileSystem) PlatformInterface {
	operator.NewOperator(fs, options.Root).InitMainOperators()

	platform := &Platform{
		options: options,
		fs:      fs,
	}

	switch options.Platform {
	case "web":
		return platform.Web()
	case "electron":
		return platform.Electron()
	default:
		panic("platform is not supported")
	}
}
