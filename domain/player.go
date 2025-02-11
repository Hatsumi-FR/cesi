package domain

type Player struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	LifePoints int     `json:"life_points"`
	Weapon     *Weapon `json:"weapon"`
}

func (p *Player) Attack(target *Player) error {
	if p.Weapon == nil {
		return attackErr
	}

	target.LifePoints -= p.Weapon.Power
	return nil
}
