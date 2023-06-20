package server

import (
	"log"
	"net/http"

	"github.com/nileshsahitya9/Blogs-Backend/internal/routes"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		router: routes.SetupRouter(),
	}
}

func (s *Server) Start(addr string) error {
	log.Printf("Server is listening on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}
