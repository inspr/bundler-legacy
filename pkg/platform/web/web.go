package web

import (
	"context"
	"fmt"
	"os"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	api "inspr.dev/primal/pkg/api"
)

type Bundler struct {
	options esbuild.BuildOptions
	root    string
	meta    api.Metadata
	mode    string
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

func NewBundler() *Bundler {
	path, _ := os.Getwd()

	return &Bundler{
		root: path,
		meta: api.NewMetadata(),

		mode: "",
		options: esbuild.BuildOptions{
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
		bundler.mode = "client"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: bundler.root,
			Sourcefile: "client.js",
			Loader:     esbuild.LoaderTSX,
		}
	case "server":
		bundler.mode = "server"
		bundler.options.Stdin = &esbuild.StdinOptions{
			Contents:   serverEntry,
			ResolveDir: bundler.root,
			Sourcefile: "server.js",
			Loader:     esbuild.LoaderTSX,
		}
	default:
		panic("target must be client or server")
	}

	return bundler
}

func (bundler *Bundler) Apply(ctx context.Context, opts api.OperatorOptions) error {
	bundler.meta.State <- api.WORKING

	select {
	case <-ctx.Done():
		return nil
	default:

		if opts.Watch {
			bundler.options.Watch = &esbuild.WatchMode{
				OnRebuild: func(result esbuild.BuildResult) {
					if len(result.Errors) > 0 {
						fmt.Printf("watch build failed: %d errors\n", len(result.Errors))
					} else {
						fmt.Printf("watch build succeeded: %d warnings\n", len(result.Warnings))
					}
				},
			}
		}

		r := esbuild.Build(bundler.options)
		errs := r.Errors

		if len(errs) != 0 {
			bundler.meta.Messages <- "failed with errors"
			for _, err := range errs {
				bundler.meta.Messages <- err.Text
			}
		} else {
			bundler.meta.Messages <- "compiled for platform web with success"
		}

		total := float32(len(r.OutputFiles))
		count := float32(1.0)

		for _, out := range r.OutputFiles {
			outFile := strings.TrimPrefix(out.Path, opts.Root)

			switch bundler.mode {
			case "server":
				outFile = strings.Replace(outFile, "stdin", "server", -1)
			default:
				outFile = strings.Replace(outFile, "stdin", "client", -1)
			}
			// outFile = strings.ToLower(outFile)

			err := opts.Files.Write(outFile, out.Contents)
			if err != nil {
				fmt.Println(err)
			}

			bundler.meta.Progress <- count / total
			count = count + 1.0
		}

		err := opts.Files.Write("/meta.json", []byte(r.Metafile))

		if err != nil {
			fmt.Println(err)
		}

		bundler.meta.State <- api.DONE
		r.Stop()

		return nil
	}
}

func (b *Bundler) Meta() api.Metadata {
	return b.meta
}
