package entities

type Genre struct {
	GenreId     int    `json:"game_id" db:"game_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
