package adapters

import "cesi/domain"

type Repositorer interface {
	CreatePlayer(p domain.Player) (int, error)
	Init() error
}
