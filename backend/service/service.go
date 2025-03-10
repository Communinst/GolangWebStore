// Describe overall entities and functions
// Get back to HANDLER_MAIN to determine the relevant routes for each endpoints
// If not done inside 24 hours, you're better be dead
// Freaking idiot

//Auth and Middleware MH
// Up handler

package service

import (
	"context"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

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
type UserServiceInterface interface {
	PostUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, userId int) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, userId int) error
	PutUserRole(ctx context.Context, userId int, roleId int) error
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
	UserServiceInterface
	// Wallet
}

func New(repo *repository.Repository) *Service {
	return &Service{
		UserServiceInterface: NewUserService(repo.UserRepo),
	}
}
