package entities

type Genre struct {
	GenreId     int    `json:"genre_id" db:"genre_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
