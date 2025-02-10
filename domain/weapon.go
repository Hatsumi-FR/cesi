package domain

type Weapon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Power       int    `json:"power"`
	IsLongRange bool   `json:"is_long_range"`
}
