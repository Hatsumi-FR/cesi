package domain

import "testing"

func TestPlayer_Attack(t *testing.T) {
	player := Player{
		ID:         1,
		Name:       "Player 1",
		LifePoints: 10,
		Weapon:     Weapon{},
	}

	target := Player{
		ID:         2,
		Name:       "Player 2",
		LifePoints: 10,
		Weapon:     Weapon{},
	}

	player.Attack(&target)

	// should lower target's life points by 1
	if target.LifePoints != 9 {
		t.Errorf("Expected target")
	}
}
