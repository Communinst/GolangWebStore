package entities

import "time"

type Game struct {
	GameId      int       `json:"game_id " db:"game_id"`
	PublisherId int       `json:"publisher_id" db:"publisher_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Releasedate time.Time `json:"release_date" db:"release_date"`
	Discount    int       `json:"discount" db:"discount"`
	GenreId     []int     `json:"genre_id" db:"genre_id"`
	DeveloperId []int     `json:"developer_id" db:"developer_id"`
}

type GameRating struct {
	GameId                 int `json:"game_id " db:"game_id "`
	RatingPercentage       int `json:"good_percentage" db:"good_percentage"`
	RatingPercentageLatest int `json:"good_percentage_latest" db:"good_percentage_latest"`
}
