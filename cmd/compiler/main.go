package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

type Operator func(fs filesystem.FileSystem)

type Primal struct {
	operators []Operator
}

func (p *Primal) Add(op ...Operator) {
	p.operators = append(p.operators, op...)
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

func (p *Primal) Run(fs filesystem.FileSystem) {
	path, _ := os.Getwd()

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
		Stdin: &esbuild.StdinOptions{
			Contents:   clientEntry,
			ResolveDir: path,
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
	writeResultsToFs(r, path, fs)

	for _, op := range p.operators {
		op(fs)
	}
}

func main() {
	path, _ := os.Getwd()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fs := filesystem.NewMemoryFs()

	p := Primal{}

	p.Add(func(fs filesystem.FileSystem) {
		var htmlTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta name="theme-color" content="white">
	<meta name="theme-color" media="(prefers-color-scheme: light)" content="white">
	<meta name="theme-color" media="(prefers-color-scheme: dark)" content="black">
	<link rel="preload" href="/entry-client.css" as="style">
	<link rel="modulepreload" href="/entry-client.js">
	<link rel="modulepreload" href="/react-dom.RT5KN4QJ.js">
    <link rel="stylesheet" href="/entry-client.css">
	<title>Primal</title>
</head>
<body>
    <div id="root"></div>
</body>
<script type="module" src="/entry-client.js" ></script>
</html>`
		fs.Write("/index.html", []byte(htmlTmpl))
	})

	p.Add(func(fs filesystem.FileSystem) {
		path = path + "/__build__"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
			os.Mkdir(path+"/assets", 0755)
		}

		for key, file := range fs.Raw() {
			// TODO: catch the error and return in an "errors" chan
			f, _ := os.Create(path + key)
			f.Write(file)
		}
	})

	p.Add(func(fs filesystem.FileSystem) {
		fmt.Println(fs)
	})

	// go Start(fs)
	p.Run(fs)

	// <-c
	// os.Exit(1)
}
