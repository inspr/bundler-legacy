// TODO: create CLI for the server

package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Server is a struct that handles routes of the rest API
type Server struct {
	mux  *http.ServeMux
	ctx  context.Context
	port string
}

// NewServer configures the server
func NewServer(ctx context.Context, port string) *Server {
	s := Server{
		mux:  http.NewServeMux(),
		ctx:  ctx,
		port: port,
	}
	s.initRoutes()
	return &s
}

// Run starts the server on the port 8000
func (s *Server) Run() {
	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":"+s.port, s.mux))
}
