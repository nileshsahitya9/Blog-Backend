package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nileshsahitya9/Blogs-Backend/internal/routes"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: routes.SetupRouter(),
	}
}

func (s *Server) Start(addr string) error {
	log.Printf("Server is listening on %s\n", addr)
	log.Printf("Route setup successfully: %p", s.router)
	return http.ListenAndServe(addr, s.router)
}
