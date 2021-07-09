package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/operator/disk"
	"inspr.dev/primal/pkg/operator/logger"
	"inspr.dev/primal/pkg/operator/server"
	"inspr.dev/primal/pkg/operator/static"
	"inspr.dev/primal/pkg/platform/web"
)

type OperatorProps struct {
	Context context.Context
	Files   filesystem.FileSystem
}

type OperatorOptions struct {
	Root string

	Watch bool

	// Environment variables
	Enviroment map[string]string
}

type Node struct {
	// Refresh chan bool
	// Ready   chan bool
	ID      string
	Depends []*Node

	Run     func(props OperatorProps, opts OperatorOptions)
	isRoot  bool
	Visited bool
}

// func (w *Node) Next() []Node {
// 	return w.Child
// }

func Traverse(tree *Node, wg *sync.WaitGroup, op func(*Node)) {
	if wg == nil {
		wg = &sync.WaitGroup{}
	}

	if tree.isRoot {
		wg.Add(1)
		defer wg.Wait()
	}

	op(tree)

	for _, subtree := range tree.Depends {
		if !subtree.Visited {
			wg.Add(1)
			subtree.Visited = true
			go Traverse(subtree, wg, op)
		}
	}

	wg.Done()
}

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
<p>Goal</p>
    <div id="root"></div>
</body>
<script type="module" src="/entry-client.js" ></script>
</html>
`

var ContentTypes = map[string]string{
	".css": "text/css; charset=UTF-8",

	".js":  "application/javascript; charset=UTF-8",
	".mjs": "application/javascript; charset=UTF-8",

	".json":   "application/json; charset=UTF-8",
	".jsonld": "application/ld+json; charset=UTF-8",

	".png":  "image/png",
	".webp": "image/webp",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".svg":  "image/svg+xml; charset=utf-8",

	".woff":  "font/woff",
	".woff2": "font/woff2",
}

func SetContentType(w http.ResponseWriter, file string) {
	ext := filepath.Ext(file)
	w.Header().Add("Content-Type", ContentTypes[ext])
}

func SetCacheDuration(w http.ResponseWriter, seconds int64) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

func main() {
	// Logger := &Node{ID: "Logger", Run: func(props OperatorProps, opts OperatorOptions) {
	// 	fmt.Println(props.Files)
	// }}

	// Server := &Node{
	// 	ID: "Server",
	// 	Run: func(props OperatorProps, opts OperatorOptions) {
	// 		go http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 			var file []byte
	// 			var err error

	// 			path := r.URL.Path[0:]

	// 			switch path {
	// 			case "/":
	// 				file, err = props.Files.Get("/index.html")
	// 			default:
	// 				file, err = props.Files.Get(path)
	// 			}

	// 			if err == nil {
	// 				SetContentType(w, path)
	// 				SetCacheDuration(w, 31536000)
	// 				w.Write(file)
	// 			} else {
	// 				w.WriteHeader(404)
	// 			}
	// 		})

	// 		fmt.Printf("Available on http://127.0.0.1:%d\n", 3049)
	// 		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3049), nil))
	// 	}}

	// Disk := &Node{
	// 	ID: "Disk",
	// 	Run: func(props OperatorProps, opts OperatorOptions) {
	// 		path := opts.Root + "/__build__"

	// 		if _, err := os.Stat(opts.Root); os.IsNotExist(err) {
	// 			os.Mkdir(path, 0755)
	// 			os.Mkdir(path+"/assets", 0755)
	// 		}

	// 		for key, file := range props.Files.Raw() {
	// 			// TODO: catch the error and return in an "errors" chan
	// 			f, _ := os.Create(path + key)
	// 			f.Write(file)
	// 		}
	// 	}, Depends: []*Node{Logger, Server}}

	// Html := &Node{ID: "Html", Run: func(props OperatorProps, opts OperatorOptions) {
	// 	props.Files.Write("/index.html", []byte(htmlTmpl))
	// }, Depends: []*Node{Disk}}

	// Bundler := &Node{
	// 	ID:      "Bundler",
	// 	Depends: []*Node{Html},
	// 	Run: func(props OperatorProps, opts OperatorOptions) {
	// 		time.Sleep(1 * time.Second)
	// 	}}

	// Root := &Node{
	// 	ID:      "Root",
	// 	Depends: []*Node{Bundler},
	// 	isRoot:  true,
	// }

	fs := filesystem.NewMemoryFs()
	root, _ := os.Getwd()

	// props := OperatorProps{Files: fs}
	// opts := OperatorOptions{Root: root}

	// Traverse(Root, nil, func(n *Node) {
	// 	if n.Run != nil {
	// 		fmt.Println(n.ID)
	// 		n.Run(props, opts)
	// 	}
	// })

	primal := NewCompiler(root, fs)

	BundlerClient := web.NewBundler().WithMinification().Target("client")
	// BundlerServer := web.NewBundler().WithMinification().Target("server")

	HtmlGen := web.NewHtml()
	Disk := disk.NewDisk("./__build__")
	Server := server.NewServer(3049)
	Logger := logger.NewLogger()
	Static := static.NewStatic([]string{"./template/sw.js"})

	primal.
		Add(Static).
		Add(BundlerClient).
		Add(HtmlGen).
		Add(Logger).
		Add(Disk, Server).
		Apply()
}
