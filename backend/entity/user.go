package entities

import "time"

type User struct {
	UserId     int       `json:"user_id" db:"user_id"`
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	Nickname   string    `json:"nickname" db:"nickname"`
	Email      string    `json:"email" db:"email"`
	SignUpDate time.Time `json:"sign_up_date" db:"sign_up_date"`
	RoleId     int       `json:"role_id"`
}
