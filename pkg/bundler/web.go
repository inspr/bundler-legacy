package bundler

import (
	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

// NewBundler returns a bundler structure with the given configs
func NewWebBundler(outdir string, fs filesystem.FileSystem) *Bundler {
	return &Bundler{
		refresh: make(chan bool, 1000),
		mode:    "client",
		outdir:  outdir,
		fs:      fs,
		options: esbuild.BuildOptions{
			Bundle:            true,
			Incremental:       true,
			Metafile:          true,
			Splitting:         true,
			Write:             false,
			ChunkNames:        "[name].[hash]",
			AssetNames:        "[name].[hash]",
			Outdir:            outdir,
			Define:            Definition,
			Loader:            LoadableExtensions,
			Platform:          esbuild.PlatformBrowser,
			Target:            esbuild.ES2015,
			LogLevel:          esbuild.LogLevelSilent,
			Sourcemap:         esbuild.SourceMapExternal,
			LegalComments:     esbuild.LegalCommentsExternal,
			Format:            esbuild.FormatESModule,
			PublicPath:        "/",
			JSXFactory:        "__jsx",
			ResolveExtensions: AddPlatformExtensions("web", Extensions),
		},
	}
}
