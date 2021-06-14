package web

import (
	"context"
	"os"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	fs "inspr.dev/primal/pkg/filesystem"
	op "inspr.dev/primal/pkg/operator"
)

type Bundler struct {
	options api.BuildOptions

	root string

	progress chan float32
	messages chan string
	done     chan bool
}

var Extensions = []string{".tsx", ".ts", ".jsx", ".js", ".wasm", ".png", ".jpg", ".svg", ".css"}
var Definition = map[string]string{"__WEB__": "true"}
var LoadableExtensions = map[string]api.Loader{
	".css": api.LoaderCSS,
	".png": api.LoaderFile,
	".svg": api.LoaderText,
}

func AddPlatformExtensions(platform string, baseExt []string) []string {
	ext := []string{}
	for _, extension := range baseExt {
		ext = append(ext, strings.Join([]string{"." + platform, extension}, ""))
	}

	ext = append(ext, baseExt...)
	return ext
}

func NewBundler() *Bundler {
	path, _ := os.Getwd()

	return &Bundler{
		progress: make(chan float32),
		messages: make(chan string),
		done:     make(chan bool),

		root: path,

		options: api.BuildOptions{
			Bundle:            true,
			Incremental:       true,
			Metafile:          true,
			Splitting:         true,
			Write:             false,
			ChunkNames:        "[name].[hash]",
			AssetNames:        "assets/[name].[hash]",
			GlobalName:        "Primal",
			Outdir:            path,
			Define:            Definition,
			Loader:            LoadableExtensions,
			Platform:          api.PlatformBrowser,
			Target:            api.ES2015,
			LogLevel:          api.LogLevelSilent,
			Sourcemap:         api.SourceMapExternal,
			LegalComments:     api.LegalCommentsExternal,
			Format:            api.FormatESModule,
			PublicPath:        "/",
			JSXFactory:        "__jsx",
			ResolveExtensions: AddPlatformExtensions("web", Extensions),
		},
	}
}

func (bundler *Bundler) WithMinification() *Bundler {
	bundler.options.MinifySyntax = true
	bundler.options.MinifyWhitespace = true
	bundler.options.MinifyIdentifiers = true
	return bundler
}

func (bundler *Bundler) WithDevelopMode() *Bundler {
	bundler.options.Define["process.env.NODE_ENV"] = "\"development\""
	return bundler
}

func (bundler *Bundler) Target(name string) *Bundler {
	clientEntry := `
		import createApp from "@primal/web/client"
		import Root from "./template"
		createApp(Root)
	`

	serverEntry := `
	import createApp from "@primal/web/server"
	import Root from "./template"
	createApp(Root)
	`

	switch name {
	case "client":
		bundler.options.Stdin = &api.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: bundler.root,
			Sourcefile: "client.js",
			Loader:     api.LoaderTSX,
		}
	case "server":
		bundler.options.Stdin = &api.StdinOptions{
			Contents:   serverEntry,
			ResolveDir: bundler.root,
			Sourcefile: "server.js",
			Loader:     api.LoaderTSX,
		}
	default:
		panic("target must be client or server")
	}

	return bundler
}

func (bundler *Bundler) Apply(ctx context.Context, spec op.Spec, fs fs.FileSystem) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		r := api.Build(bundler.options)
		errs := r.Errors

		if len(errs) != 0 {
			bundler.messages <- "failed with errors"
			for _, err := range errs {
				bundler.messages <- err.Text
			}
		} else {
			bundler.messages <- "compiled for platform web with success"
		}

		total := float32(len(r.OutputFiles))
		count := float32(1.0)

		for _, out := range r.OutputFiles {
			outFile := strings.TrimPrefix(out.Path, spec.Root)
			outFile = strings.Replace(outFile, "stdin", "client", -1)
			// outFile = strings.ToLower(outFile)

			err := fs.Write(outFile, out.Contents)
			if err != nil {
				panic(err)
			}

			bundler.progress <- count / total
			count = count + 1.0
		}

		err := fs.Write("/meta.json", []byte(r.Metafile))
		if err != nil {
			panic(err)
		}

		bundler.done <- true

		return nil
	}
}

func (bundler *Bundler) Progress() <-chan float32 {
	return bundler.progress
}

func (bundler *Bundler) Messages() <-chan string {
	return bundler.messages
}

func (bundler *Bundler) Done() <-chan bool {
	return bundler.done
}
