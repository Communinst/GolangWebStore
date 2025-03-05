package entities

type Cart struct {
	CartId int `json:"cart_id" db:"cart_id"`
	UserId int `json:"user_id" db:"user_id"`
}

type CartGames struct {
	CartGamesId int `json:"cart_games_id" db:"cart_games_id"`
	CartId      int `json:"cart_id" db:"cart_id"`
	GameId      int `json:"game_id" db:"game_id"`
}
