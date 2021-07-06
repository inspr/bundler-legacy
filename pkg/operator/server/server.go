package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"inspr.dev/primal/pkg/api"
)

type Server struct {
	port int
	meta api.Metadata
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
		meta: api.NewMetadata(),
	}
}

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

func (s *Server) Apply(props api.OperatorProps, opts api.OperatorOptions) {

	go http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var file []byte
		var err error

		path := r.URL.Path[0:]
		s.meta.Messages <- fmt.Sprintf("serving %s", path)

		switch path {
		case "/":
			file, err = props.Files.Get("/index.html")
		default:
			file, err = props.Files.Get(path)
		}

		if err == nil {
			SetContentType(w, path)
			SetCacheDuration(w, 31536000)
			w.Write(file)
		} else {
			w.WriteHeader(404)
		}
	})

	s.meta.Messages <- fmt.Sprintf("Available on http://127.0.0.1:%d", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))

	s.meta.Done <- true

Main:
	for {
		select {
		case <-s.meta.Close:
			break Main

		case <-s.meta.Refresh:
			fmt.Println("refreshed")
		}
	}
}

func (s *Server) Meta() api.Metadata {
	return s.meta
}
