package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/mjibson/go-dsp/fft"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

type Meta struct {
	Links []string
}

func main() {
	f, err := os.ReadFile("./lib/web/server/index.html")
	if err != nil {
		panic(err)
	}

	meta := Meta{
		[]string{
			"teste.js",
			"teste.js",
			"teste.js",
			"teste.js",
			"teste.js",
		},
	}

	tmpl, err := template.New("test").Parse(string(f))

	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(nil)

	err = tmpl.Execute(buff, meta)
	if err != nil {
		panic(err)
	}

	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	b, err := m.Bytes("text/html", buff.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	fmt.Println(fft.FFTReal([]float64{1, 2, 3}))
}
