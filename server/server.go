package server

import (
	"cesi/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Server struct {
	repository repository.Repositorer
}

func NewServer(r repository.Repositorer) *Server {
	return &Server{
		repository: r,
	}
}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/players", s.createPlayerHandler)
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
