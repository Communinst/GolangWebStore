package entities

import "time"

type Game struct {
	GameId      int       `json:"game_id" db:"game_id"`
	PublisherId int       `json:"publisher_id" db:"publisher_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Releasedate time.Time `json:"release_date" db:"release_date"`
	Rating      float32   `json:"rating" db:"rating"`
	Discount    int       `json:"discount" db:"discount"`
	GenreId     []int     `json:"genre_id" db:"genre_id"`
}

