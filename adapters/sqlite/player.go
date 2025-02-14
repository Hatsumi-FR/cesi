package sqlite

import (
	"cesi/domain"
)

func (r *Repository) CreatePlayer(p domain.Player) (int, error) {
	request := "INSERT INTO players (name, life_points) VALUES (?, ?)"
	result, err := r.db.Exec(request, p.Name, p.LifePoints)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
