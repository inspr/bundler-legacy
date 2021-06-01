package bundler

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

type Bundler struct {
	options api.BuildOptions
}

var clientEntry = `
import "./primal.ts"
import "./primal.css"
import ReactDOM from 'react-dom'
import React from 'react'
import Root from "./index.tsx"
	
ReactDOM.render(
	<Root/>,
	document.getElementById('root')
)
`

var serverEntry = `
import {jsx as __jsx} from 'react/jsx-runtime';

import React from 'react'
import "./primal.ts"
import Root from "./index.tsx"
import ReactDOMServer from 'react-dom/server'

const code = ReactDOMServer.renderToString(<Root/>)
run(code)
`

func NewBundler() *Bundler {
	path, _ := os.Getwd()

	return &Bundler{
		options: api.BuildOptions{
			// EntryPoints: []string{path + "/input.tsx"},
			// Outfile:     "bundle.js",
			// Outdir:     path + "/__build__",
			Platform: api.PlatformBrowser,
			Bundle:   true,
			Write:    false,
			// AssetNames: "assets/[name]-[hash]",
			// EntryNames: "[dir]/[name]-[hash]",
			ChunkNames: "[name].[hash]",
			Splitting:  false,
			// Outfile: path + "/main.js",
			Outdir:        path + "/__build__",
			Define:        map[string]string{"__WEB__": "true"},
			Metafile:      true,
			Target:        api.ES2015,
			LogLevel:      api.LogLevelSilent,
			Sourcemap:     api.SourceMapExternal,
			LegalComments: api.LegalCommentsExternal,
			JSXFactory:    "__jsx",
			Loader: map[string]api.Loader{
				".css": api.LoaderCSS,
				".png": api.LoaderFile,
				".svg": api.LoaderText,
			},
			PublicPath:        "/",
			Format:            api.FormatESModule,
			Incremental:       true,
			ResolveExtensions: []string{".web.tsx", ".web.ts", ".ts", ".js", ".tsx", ".jsx", ".png"},
		},
	}
}

func (b *Bundler) WithMinification() *Bundler {
	b.options.MinifySyntax = true
	b.options.MinifyWhitespace = true
	b.options.MinifyIdentifiers = true
	return b
}

func (b *Bundler) WithDevelopMode() *Bundler {
	b.options.Define["process.env.NODE_ENV"] = "\"development\""
	return b
}

func (b *Bundler) AsServer() *Bundler {
	path, _ := os.Getwd()

	b.options.Stdin = &api.StdinOptions{
		Contents:   serverEntry,
		ResolveDir: path,
		Sourcefile: "server.js",
		Loader:     api.LoaderTSX,
	}

	return b
}

func (b *Bundler) AsClient() *Bundler {
	path, _ := os.Getwd()

	b.options.Stdin = &api.StdinOptions{
		Contents:   clientEntry,
		ResolveDir: path,
		Sourcefile: "client.js",
		Loader:     api.LoaderTSX,
	}

	return b
}

func (b *Bundler) Build() api.BuildResult {
	r := api.Build(b.options)
	return r
}
