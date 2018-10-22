package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../api/root"
	"../router"
	"github.com/gorilla/mux"
)

// Server represents a server instance.
type Server struct {
	router *mux.Router
}

// configureRouter configures the router for the Server instance.
func (s *Server) configureRouter() {
	var appRoutes = make(router.Routes, 0)
	appRoutes = append(appRoutes, root.Routes...)

	s.router = router.BuildRouter(appRoutes)
}

// Start initializes the web server and starts listening on a port.
func (s *Server) Start() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	s.configureRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}
