package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../api/root"
	"../database"
	"../router"
	"github.com/gorilla/mux"
)

// Server represents a server instance.
type Server struct {
	router *mux.Router
}

// configureRouter configures the router for the Server instance.
func (s *Server) configureRouter() error {
	var appRoutes = make(router.Routes, 0)
	appRoutes = append(appRoutes, root.Routes...)

	var db database.DatabaseContract = &database.InMemory{}
	err := db.Connect(nil)
	if err != nil {
		return err
	}

	s.router = router.BuildRouter(appRoutes, db)
	return nil
}

// Start initializes the web server and starts listening on a port.
func (s *Server) Start() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	if err := s.configureRouter(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}
