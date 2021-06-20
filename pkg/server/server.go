package server

import (
	"context"
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

func (s *Server) Apply(ctx context.Context, opts api.OperatorOptions) error {
	s.meta.State <- api.WORKING

	go http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var file []byte
		var err error

		path := r.URL.Path[0:]

		switch path {
		case "/":
			file, err = opts.Files.Get("/index.html")
		default:
			file, err = opts.Files.Get(path)
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
	// s.done <- true
	s.meta.Progress <- 1.0
	s.meta.State <- api.READY

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))
	return nil
}

func (s *Server) Meta() api.Metadata {
	return s.meta
}
