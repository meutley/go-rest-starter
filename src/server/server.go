package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../router"
	"github.com/gorilla/mux"
)

// Server represents a server instance.
type Server struct {
	router mux.Route
}

// Start initializes the web server and starts listening on a port.
func (s *Server) Start() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router := router.BuildRouter(router.AppRoutes)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
