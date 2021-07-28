package platform

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	esbuild "github.com/evanw/esbuild/pkg/api"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator"
	"inspr.dev/primal/pkg/workflow"
)

type Web struct {
	*Platform

	workflow workflow.Workflow
}

func (p *Platform) Web() PlatformInterface {
	web := &Web{
		Platform: p,
	}

	for _, ops := range operator.MainOps {
		web.workflow.Add(ops.Task())
	}

	return web
}

var Extensions = []string{".tsx", ".ts", ".jsx", ".js", ".wasm", ".png", ".jpg", ".svg", ".css"}
var Definition = map[string]string{"__WEB__": "true"}
var LoadableExtensions = map[string]esbuild.Loader{
	".css": esbuild.LoaderCSS,
	".png": esbuild.LoaderFile,
	".svg": esbuild.LoaderText,
}

func AddPlatformExtensions(platform string, baseExt []string) []string {
	ext := []string{}
	for _, extension := range baseExt {
		ext = append(ext, strings.Join([]string{"." + platform, extension}, ""))
	}

	ext = append(ext, baseExt...)
	return ext
}

func writeResultsToFs(r esbuild.BuildResult, path string, fs filesystem.FileSystem) {
	for _, out := range r.OutputFiles {
		outFile := strings.TrimPrefix(out.Path, path)
		outFile = strings.Replace(outFile, "stdin", "entry-client", -1)

		err := fs.Write(outFile, out.Contents)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (w *Web) Run() {
	clientEntry := fmt.Sprintf(`
	import createApp from "@primal/web/client"
	import Root from "%s"
	createApp(Root)
	`, "./template")

	options := esbuild.BuildOptions{
		Bundle:            true,
		Incremental:       true,
		Metafile:          true,
		Splitting:         true,
		Write:             false,
		ChunkNames:        "[name].[hash]",
		AssetNames:        "[name].[hash]",
		Outdir:            w.options.Root,
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
		Stdin: &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: w.options.Root,
			Sourcefile: "client.js",
			Loader:     esbuild.LoaderTSX,
		},
		// Watch: &esbuild.WatchMode{
		// 	OnRebuild: func(r esbuild.BuildResult) {
		// 		if len(r.Errors) > 0 {
		// 			fmt.Printf("watch build failed: %d errors\n", len(r.Errors))
		// 		} else {
		// 			fmt.Printf("watch build succeeded: %d warnings\n", len(r.Warnings))
		// 		}

		// 		fs.Clean()

		// 		writeResultsToFs(r, path, fs)
		// 		for _, op := range p.operators {
		// 			op(fs)
		// 		}
		// 	},
		// },
	}

	options.MinifySyntax = true
	options.MinifyWhitespace = true
	options.MinifyIdentifiers = true

	r := esbuild.Build(options)
	writeResultsToFs(r, w.options.Root, w.fs)

	w.workflow.Run()
}

func GracefullShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	msg := <-c
	fmt.Println(msg)
	os.Exit(1)
}

func (w *Web) Watch() {
	clientEntry := fmt.Sprintf(`
	import createApp from "@primal/web/client"
	import Root from "%s"
	createApp(Root)
	`, "./template")

	options := esbuild.BuildOptions{
		Bundle:            true,
		Incremental:       true,
		Metafile:          true,
		Splitting:         true,
		Write:             false,
		ChunkNames:        "[name].[hash]",
		AssetNames:        "[name].[hash]",
		Outdir:            w.options.Root,
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
		Stdin: &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: w.options.Root,
			Sourcefile: "client.js",
			Loader:     esbuild.LoaderTSX,
		},
		Watch: &esbuild.WatchMode{
			OnRebuild: func(r esbuild.BuildResult) {
				if len(r.Errors) > 0 {
					fmt.Printf("watch build failed: %d errors\n", len(r.Errors))
				} else {
					fmt.Printf("watch build succeeded: %d warnings\n", len(r.Warnings))
				}

				w.fs.Clean()

				writeResultsToFs(r, w.options.Root, w.fs)

				w.workflow.Run()
			},
		},
	}

	options.MinifySyntax = true
	options.MinifyWhitespace = true
	options.MinifyIdentifiers = true

	r := esbuild.Build(options)
	writeResultsToFs(r, w.options.Root, w.fs)

	w.workflow.Run()

	go Start(w.fs)
	GracefullShutdown()
}
