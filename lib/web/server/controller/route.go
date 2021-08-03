package controller

import (
	"inspr.dev/primal/lib/web/server/handlers"
)

// initRoutes defines the server initialization method, which will add all
// the possible routes the server will handle
func (s *Server) initRoutes() {
	h := handlers.NewHandler(s.ctx)

	s.mux.HandleFunc("/start", h.InitBuildDir())

	s.mux.HandleFunc("/", h.ServeFiles())

	s.mux.HandleFunc("/healtz", h.HealthCheck())
}
