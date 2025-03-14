package entities

import "time"

type Ownership struct {
	OwnershipId  int       `json:"ownership_id" db:"ownership_id"`
	UserId       int       `json:"user_id" db:"user_id"`
	GameId       int       `json:"game_id" db:"game_id"`
	MinutesSpent int64     `json:"minutes_spent" db:"minutes_spent"`
	ReceiptDate  time.Time `json:"receipt_date" db:"receipt_date"`
	//AchievedId   []int     `json:"achievement_id"`
}
