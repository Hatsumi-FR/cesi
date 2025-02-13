package domain

import (
	"errors"
	"testing"
)

func TestPlayer_Attack(t *testing.T) {
	tests := []struct {
		name      string
		p1        *Player
		p2        *Player
		wantedErr error
	}{
		{
			name: "should lose life points",
			p1: &Player{
				ID:         1,
				Name:       "player 1",
				LifePoints: 10,
				Weapon: &Weapon{
					ID:          1,
					Name:        "weapon 1",
					Power:       1,
					IsLongRange: false,
				},
			},
			p2: &Player{
				ID:         2,
				Name:       "weapon 2",
				LifePoints: 10,
				Weapon: &Weapon{
					ID:          1,
					Name:        "weapon 1",
					Power:       1,
					IsLongRange: false,
				},
			},
			wantedErr: nil,
		},
		{
			name: "should return an error because weapon doesn't exist",
			p1: &Player{
				ID:         1,
				Name:       "player 1",
				LifePoints: 10,
				Weapon:     nil,
			},
			p2: &Player{
				ID:         2,
				Name:       "weapon 2",
				LifePoints: 10,
				Weapon:     nil,
			},
			wantedErr: attackErr,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.p1.Attack(test.p2)
			if !errors.Is(err, test.wantedErr) {
				t.Error(err.Error())
			}
		})
	}
}

func TestPlayer_Heal(t *testing.T) {
	// init player
	p1 := Player{
		ID:         1,
		Name:       "p1",
		LifePoints: 10,
		Weapon:     nil,
	}
	t.Run(t.Name(), func(t *testing.T) {
		p1.Heal()
		if p1.LifePoints != 11 {
			t.Errorf("player should have 11 hp \n current hp : %d", p1.LifePoints)
		}
	})
}
