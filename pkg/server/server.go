package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"inspr.dev/primal/pkg/filesystem"
)

func SetContentType(w http.ResponseWriter, file string) {
	ext := filepath.Ext(file)

	switch ext {
	case ".css":
		w.Header().Add("Content-Type", "text/css; charset=UTF-8")
	case ".js":
		w.Header().Add("Content-Type", "application/javascript; charset=UTF-8")
	case ".json":
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	case ".png":
		w.Header().Add("Content-Type", "image/png; charset=utf-8")
	}
}

func SetCacheDuration(w http.ResponseWriter, seconds int64) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

func Run(files filesystem.FileSystem) {
	go http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var file []byte
		var err error

		path := r.URL.Path[0:]

		switch path {
		case "/":
			file, err = files.Get("/index.html")
		default:
			file, err = files.Get(path)
		}

		if err == nil {
			SetContentType(w, path)
			SetCacheDuration(w, 31536000)
			w.Write(file)
		} else {
			w.WriteHeader(404)
		}
	})

	fmt.Println("Available on http://127.0.0.1:3049")
	log.Fatal(http.ListenAndServe(":3049", nil))
}
