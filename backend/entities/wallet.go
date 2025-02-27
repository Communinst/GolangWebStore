package entities

type Wallet struct {
	UserId  int `json:"user_id" db:"user_id"`
	Balance int `json:"balance" db:"balance"`
}
