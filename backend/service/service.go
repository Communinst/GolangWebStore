// Describe overall entities and functions
// Get back to HANDLER_MAIN to determine the relevant routes for each endpoints
// If not done inside 24 hours, you're better be dead
// Freaking idiot
package service

import "github.com/Communinst/GolangWebStore/backend/repository"

type Authorization interface {
}
type Cart interface {
}
type Company interface {
}
type Discount interface {
}
type Game interface {
}
type Genre interface {
}
type Ownership interface {
}
type Review interface {
}
type Role interface {
}
type User interface {
}
type Wallet interface {
}

type Service struct {
	// Authorization
	// Cart
	//Company
	// Discount
	//Game
	// Genre
	// Ownership
	// Review
	// Role
	User
	// Wallet
}

func New(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.UserRepo),
	}
}
