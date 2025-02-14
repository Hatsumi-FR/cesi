package inmemory

import (
	"cesi/domain"
)

type Repository struct {
	players []domain.Player
}

func (r *Repository) CreatePlayer(p domain.Player) (int, error) {
	r.players = append(r.players, p)
	return len(r.players) - 1, nil
}

func (r *Repository) Init() error {
	return nil
}
