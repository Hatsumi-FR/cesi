package server

import (
	"cesi/domain"
	"encoding/json"
	"net/http"
)

type createPlayerResponse struct {
	ID int `json:"id"`
}

func (s *Server) CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	p := domain.Player{}
	// unmarshal the request body into the player struct
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// create the player
	id, err := s.repository.CreatePlayer(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := createPlayerResponse{
		ID: id,
	}
	// marshal the player struct into json
	respEncoded, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respEncoded)
}
