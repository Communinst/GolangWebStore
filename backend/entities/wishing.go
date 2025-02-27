package entities

type Wishing struct {
	WishingId int `json:"wishing_id" db:"wishing_id"`
	UserId    int `json:"user_id" db:"user_id"`
	GameId    int `json:"game_id" db:"game_id"`
}
