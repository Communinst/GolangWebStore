package entities

import "time"

type Review struct {
	ReviewId    int       `json:"review_id" db:"review_id"`
	Recommended bool      `json:"recommended" db:"recommended"`
	Message     string    `json:"message" db:"message"`
	UserId      int       `json:"user_id" db:"user_id"`
	GameId      int       `json:"game_id" db:"game_id"`
	Date        time.Time `json:"date" db:"date"`
}
