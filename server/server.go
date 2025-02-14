package server

import (
	"cesi/adapters"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type Server struct {
	repository adapters.Repositorer
}

func NewServer(r adapters.Repositorer) *Server {
	return &Server{
		repository: r,
	}
}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/players", s.CreatePlayerHandler)
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
