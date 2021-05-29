package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"inspr.dev/primal/pkg/bundler"
)

func main() {
	path, _ := os.Getwd()
	files := make(map[string][]byte)

	b := bundler.NewBundler()
	result := b.
		// WithMinification().
		WithDevelopMode().
		AsClient().
		// AsServer().
		Build()

	if len(result.Errors) > 0 {
		fmt.Println(result.Errors)
		os.Exit(1)
	}

	info := make(map[string]float64)

	for _, out := range result.OutputFiles {
		outFile := strings.TrimPrefix(out.Path, path)
		files[outFile] = out.Contents

		var buff bytes.Buffer
		zw := gzip.NewWriter(&buff)

		_, err := zw.Write(out.Contents)
		if err != nil {
			log.Fatal(err)
		}

		writeAs := strings.Replace(out.Path, "stdin.js", "main.js", -1)
		outName := strings.Replace(outFile, "stdin.js", "main.js", -1)
		info[outName] = float64(len(out.Contents)) / 1024.0

		// Write to file
		file, err := os.Create(writeAs)
		if err != nil {
			log.Fatal(err)
		}

		file.Write(out.Contents)
		defer file.Close()
	}

	for key, val := range info {
		key := strings.TrimPrefix(key, "/__build__/")
		fmt.Printf("%s %s %.2f%s\n", aurora.BrightGreen("✔︎"), key, aurora.Bold(val), aurora.Bold("KB"))
	}

	// vm.Run(files)

	// 	test := `
	// 	import {jsx, jsxs} from 'react/jsx-runtime'
	// import "./test.css"

	// // implement a fix for react jsx new format as defined by react 17
	// // the order of the elements is different there and the key is external
	// globalThis.__jsx = function (type, props, ...children) {
	//     if (typeof props === "undefined" ||  !props) {
	// 		props = {}
	// 	}

	//     let {key, ...otherProps} = props

	// 	if (children.length === 0) {
	// 		children = null
	//         return jsx(type, {...otherProps, children}, key)
	//     } else {
	//         return jsxs(type, {...otherProps, children}, key)
	//     }

	// }
	// 	`

	// result := api.Transform(test, api.TransformOptions{
	// 	// JSXFactory: "__pjsx",
	// 	MinifySyntax:      true,
	// 	MinifyWhitespace:  true,
	// 	MinifyIdentifiers: true,
	// 	Loader:            api.LoaderTSX,
	// })

	// fmt.Println(string(result.Code))
}
