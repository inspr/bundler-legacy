// TODO: create CLI for the server

package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"inspr.dev/primal/lib/web/server/vm"
)

// Server is a struct that handles routes of the rest API
type Server struct {
	mux  *http.ServeMux
	ctx  context.Context
	port string
	path string

	machine vm.Interface
}

// NewServer configures the server
func NewServer(ctx context.Context, port, path string, machine vm.Interface) *Server {
	s := Server{
		mux:     http.NewServeMux(),
		ctx:     ctx,
		port:    port,
		path:    path,
		machine: machine,
	}
	s.initRoutes()
	return &s
}

// Run starts the server on the port 8000
func (s *Server) Run() {
	fmt.Println("server running on port: ", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, s.mux))
}
