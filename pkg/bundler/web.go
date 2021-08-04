package bundler

import (
	"fmt"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

// NewBundler returns a bundler structure with the given configs.
// If wMode (watch mode) is true, the bundler should be on client mode.
// Otherwise, it will be on server mode
func NewWebBundler(wMode bool, outdir string, fs filesystem.FileSystem) *Bundler {
	bundler := Bundler{
		refresh: make(chan bool, 1000),
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
			PublicPath:        "/static/",
			JSXFactory:        "__jsx",
			ResolveExtensions: AddPlatformExtensions("web", Extensions),
		},
	}

	if wMode {
		clientEntry := fmt.Sprintf(`
		import createApp from "@primal/web/client"
		import Root from "%s"
		createApp(Root)
		`, "./template")

		bundler.mode = "client"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: bundler.outdir,
			Sourcefile: "client.js",
			Loader:     esbuild.LoaderTSX,
		}
	} else {
		serverEntry := fmt.Sprintf(`
		import createApp from "@primal/web/server"
		import Root from "%s"
		createApp(Root)
		`, "./template")

		bundler.mode = "server"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   serverEntry,
			ResolveDir: bundler.outdir,
			Sourcefile: "server.js",
			Loader:     esbuild.LoaderTSX,
		}
	}

	return &bundler
}
