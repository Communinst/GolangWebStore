package entities

type Wallet struct {
	WalletId int `json:"wallet_id" db:"wallet_id"`
	UserId   int `json:"user_id" db:"user_id"`
	Balance  int `json:"balance" db:"balance"`
}
