package bundler

import (
	"fmt"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

// Bundler contains the data necessary for a bundler to run
type Bundler struct {
	Mode    string
	Outdir  string
	Fs      filesystem.FileSystem
	Options esbuild.BuildOptions
	Refresh chan bool
}

// Extensions define the main supported extension types
var Extensions = []string{".tsx", ".ts", ".jsx", ".js", ".wasm", ".png", ".jpg", ".svg", ".css"}

// TODO: Definition should be dynamicly setted depending on platform
var Definition = map[string]string{"__WEB__": "true", "process.env.NODE_ENV": "'production'"}
var LoadableExtensions = map[string]esbuild.Loader{
	".css": esbuild.LoaderCSS,
	".png": esbuild.LoaderFile,
	".svg": esbuild.LoaderText,
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
	r := esbuild.Build(bundler.Options)
	bundler.writeResultsToFs(r)
}

// Watch runs Bundler in watch mode
func (bundler *Bundler) Watch() {
	bundler.Options.Watch = &esbuild.WatchMode{
		OnRebuild: func(r esbuild.BuildResult) {
			if len(r.Errors) > 0 {
				fmt.Printf("watch build failed: %d errors\n", len(r.Errors))
			} else {
				fmt.Printf("watch build succeeded: %d warnings\n", len(r.Warnings))
			}

			bundler.writeResultsToFs(r)
			bundler.Refresh <- true
		},
	}

	r := esbuild.Build(bundler.Options)
	bundler.writeResultsToFs(r)
}

// WithMinification sets bundler to run with minification options
func (bundler *Bundler) WithMinification() *Bundler {
	bundler.Options.MinifySyntax = true
	bundler.Options.MinifyWhitespace = true
	bundler.Options.MinifyIdentifiers = true
	return bundler
}

// WithDevelopMode makes bundler run in development mode
func (bundler *Bundler) WithDevelopMode() *Bundler {
	bundler.Options.Define["process.env.NODE_ENV"] = "\"development\""

	return bundler
}

func (bundler *Bundler) writeResultsToFs(r esbuild.BuildResult) {
	for _, out := range r.OutputFiles {
		outFile := strings.TrimPrefix(out.Path, bundler.Outdir)

		switch bundler.Mode {
		case "server":
			outFile = strings.Replace(outFile, "stdin", "entry-server", -1)
		default:
			outFile = strings.Replace(outFile, "stdin", "entry-client", -1)
		}

		outFile = "/static" + outFile

		err := bundler.Fs.Write(outFile, out.Contents)
		if err != nil {
			fmt.Println(err)
		}
	}
}
