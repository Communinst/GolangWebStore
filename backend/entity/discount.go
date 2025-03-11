package entities

import "time"

type Discount struct {
	DiscountId    int       `json:"discount_id" db:"discount_id"`
	GameId        int       `json:"game_id" db:"game_id"`
	DiscountValue int       `json:"discount_value" db:"discount_value"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	CeaseDate     time.Time `json:"cease_date" db:"cease_date"`
}



