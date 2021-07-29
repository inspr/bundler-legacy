package bundler

import (
	"fmt"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

// Bundler contains the data necessary for a bundler to run
type Bundler struct {
	mode    string
	outdir  string
	fs      filesystem.FileSystem
	options esbuild.BuildOptions
	refresh chan bool
}

// NewBundler returns a bundler structure with the given configs
func NewBundler(outdir string, fs filesystem.FileSystem) *Bundler {
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

// Extensions define the main supported extension types
var Extensions = []string{".tsx", ".ts", ".jsx", ".js", ".wasm", ".png", ".jpg", ".svg", ".css"}

// TODO: Definition should be dynamicly setted depending on platform
var Definition = map[string]string{"__WEB__": "true"}
var LoadableExtensions = map[string]esbuild.Loader{
	".css": esbuild.LoaderCSS,
	".png": esbuild.LoaderFile,
	".svg": esbuild.LoaderText,
}

func (bundler *Bundler) Refresh() chan bool {
	return bundler.refresh
}

// TODO: review this method
// AddPlatformExtensions adds given extentions to support the given platform
func AddPlatformExtensions(platform string, baseExt []string) []string {
	ext := []string{}
	for _, extension := range baseExt {
		ext = append(ext, strings.Join([]string{"." + platform, extension}, ""))
	}

	ext = append(ext, baseExt...)
	return ext
}

// Build builds the files in the filesystem
func (bundler *Bundler) Build() {
	r := esbuild.Build(bundler.options)
	bundler.writeResultsToFs(r)
}

// Watch runs Bundler in watch mode
func (bundler *Bundler) Watch() {
	bundler.options.Watch = &esbuild.WatchMode{
		OnRebuild: func(r esbuild.BuildResult) {
			fmt.Println("REBUILDING")
			if len(r.Errors) > 0 {
				fmt.Printf("watch build failed: %d errors\n", len(r.Errors))
			} else {
				fmt.Printf("watch build succeeded: %d warnings\n", len(r.Warnings))
			}

			bundler.writeResultsToFs(r)
			bundler.refresh <- true
		},
	}

	r := esbuild.Build(bundler.options)
	bundler.writeResultsToFs(r)
}

// WithMinification sets bundler to run with minification options
func (bundler *Bundler) WithMinification() *Bundler {
	bundler.options.MinifySyntax = true
	bundler.options.MinifyWhitespace = true
	bundler.options.MinifyIdentifiers = true
	return bundler
}

// WithDevelopMode makes bundler run in development mode
func (bundler *Bundler) WithDevelopMode() *Bundler {
	bundler.options.Define["process.env.NODE_ENV"] = "\"development\""

	return bundler
}

// TODO: Fix root to start from ./template
// Target sets bundler to work in client or server mode
func (bundler *Bundler) Target(name string) *Bundler {
	clientEntry := fmt.Sprintf(`
	import createApp from "@primal/web/client"
	import Root from "%s"
	createApp(Root)
	`, "./template")

	serverEntry := fmt.Sprintf(`
	import createApp from "@primal/web/server"
	import Root from "%s"
	createApp(Root)
	`, "./template")

	switch name {
	case "client":
		bundler.mode = "client"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: bundler.outdir,
			Sourcefile: "client.js",
			Loader:     esbuild.LoaderTSX,
		}
	case "server":
		bundler.mode = "server"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   serverEntry,
			ResolveDir: bundler.outdir,
			Sourcefile: "server.js",
			Loader:     esbuild.LoaderTSX,
		}
	default:
		panic("target must be client or server")
	}

	return bundler
}

func (bundler *Bundler) writeResultsToFs(r esbuild.BuildResult) {
	for _, out := range r.OutputFiles {
		outFile := strings.TrimPrefix(out.Path, bundler.outdir)

		switch bundler.mode {
		case "server":
			outFile = strings.Replace(outFile, "stdin", "entry-server", -1)
		default:
			outFile = strings.Replace(outFile, "stdin", "entry-client", -1)
		}

		err := bundler.fs.Write(outFile, out.Contents)
		if err != nil {
			fmt.Println(err)
		}
	}
}
