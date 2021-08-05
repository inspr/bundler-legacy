package platform

import (
	"context"
	"fmt"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/bundler"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

// Web defines a web platform data
type Web struct {
	*Platform
}

// Web returns a web platform with it's tasks
func (p *Platform) Web() api.PlatformInterface {
	web := &Web{
		Platform: p,
	}
	for _, ops := range operator.MainOps {
		web.Platform.Workflow.Add(ops.Task())
	}

	addDependencies(web.Platform.Workflow.Tasks)

	return web
}

// Run executes the workflow for the web platform
func (w *Web) Run() {
	clientBundler := newClientBundler(w.Options.Root, w.Fs)
	serverBunder := newServerBundler(w.Options.Root, w.Fs)

	clientBundler.Build()
	serverBunder.Build()

	w.Platform.Workflow.Run(context.WithCancel(context.Background()))
}

// Watch executes the workflow for the web platform in watch mode
func (w *Web) Watch(ctx context.Context, cancel context.CancelFunc) {
	clientBundler := newClientBundler(w.Options.Root, w.Fs)
	clientBundler.Watch()

	w.Platform.Workflow.Run(ctx, cancel)

	server := Server{
		reload: clientBundler.Refresh,
	}

	server.Start(ctx, w.Fs)
}

// newClientBundler returns a client type bundler structure with the given configs
func newClientBundler(outdir string, fs filesystem.FileSystem) *bundler.Bundler {
	clientEntry := fmt.Sprintf(`
		import createApp from "@primal/web/client"
		import Root from "%s"
		createApp(Root)
		`, "./template")

	return &bundler.Bundler{
		Mode:    "client",
		Refresh: make(chan bool, 1000),
		Outdir:  outdir,
		Fs:      fs,
		Options: esbuild.BuildOptions{
			Bundle:            true,
			Incremental:       true,
			Metafile:          true,
			Splitting:         true,
			Write:             false,
			ChunkNames:        "[name].[hash]",
			AssetNames:        "[name].[hash]",
			Outdir:            outdir,
			Define:            bundler.Definition,
			Loader:            bundler.LoadableExtensions,
			Platform:          esbuild.PlatformBrowser,
			Target:            esbuild.ES2015,
			LogLevel:          esbuild.LogLevelSilent,
			Sourcemap:         esbuild.SourceMapExternal,
			LegalComments:     esbuild.LegalCommentsExternal,
			Format:            esbuild.FormatESModule,
			PublicPath:        "/static/",
			JSXFactory:        "__jsx",
			ResolveExtensions: bundler.AddPlatformExtensions("web", bundler.Extensions),
			Stdin: &esbuild.StdinOptions{
				Contents:   clientEntry,
				ResolveDir: outdir,
				Sourcefile: "client.js",
				Loader:     esbuild.LoaderTSX,
			},
		},
	}
}

// newServerBundler returns a server type bundler structure with the given configs
func newServerBundler(outdir string, fs filesystem.FileSystem) *bundler.Bundler {
	serverEntry := fmt.Sprintf(`
		import createApp from "@primal/web/server"
		import Root from "%s"
		createApp(Root)
		`, "./template")

	return &bundler.Bundler{
		Mode:    "server",
		Refresh: make(chan bool, 1000),
		Outdir:  outdir,
		Fs:      fs,
		Options: esbuild.BuildOptions{
			Bundle:            true,
			Incremental:       true,
			Metafile:          true,
			Splitting:         false,
			Write:             false,
			ChunkNames:        "[name].[hash]",
			AssetNames:        "[name].[hash]",
			Outdir:            outdir,
			Define:            bundler.Definition,
			Loader:            bundler.LoadableExtensions,
			Platform:          esbuild.PlatformNode,
			Target:            esbuild.ES2015,
			LogLevel:          esbuild.LogLevelSilent,
			Sourcemap:         esbuild.SourceMapExternal,
			LegalComments:     esbuild.LegalCommentsExternal,
			Format:            esbuild.FormatCommonJS,
			PublicPath:        "/static/",
			JSXFactory:        "__jsx",
			ResolveExtensions: bundler.AddPlatformExtensions("web", bundler.Extensions),
			Stdin: &esbuild.StdinOptions{
				Contents:   serverEntry,
				ResolveDir: outdir,
				Sourcefile: "server.js",
				Loader:     esbuild.LoaderTSX,
			},
		},
	}
}

func addDependencies(tasks map[string]*workflow.Task) {
	tasks["disk"].DependsOn(tasks["html"])
	tasks["logger"].DependsOn(tasks["disk"])
}
