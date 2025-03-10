// Describe overall entities and functions
// Get back to HANDLER_MAIN to determine the relevant routes for each endpoints
// If not done inside 24 hours, you're better be dead
// Freaking idiot

// Up handler

package service

import (
	"context"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type Cart interface {
}

type Discount interface {
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
type AuthServiceInterface interface {
	GenerateAuthToken(user *entities.User, secret string, expireTime int) (string, error)
	PostUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, userId int) (*entities.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*entities.User, error)
}
type Wallet interface {
}

type CompanyServiceInterface interface {
	PostCompany(ctx context.Context, company *entities.Company) error
	GetCompany(ctx context.Context, companyId int) (*entities.Company, error)
	GetAllCompanies(ctx context.Context) ([]entities.Company, error)
	DeleteCompany(ctx context.Context, companyId int) error
}
type GameServiceInterface interface {
	PostGame(ctx context.Context, game *entities.Game) error
	GetGame(ctx context.Context, gameId int) (*entities.Game, error)
	GetAllGames(ctx context.Context) ([]entities.Game, error)
	DeleteGame(ctx context.Context, gameId int) error
	PutGamePrice(ctx context.Context, gameId int, newPrice int) error
}

type Service struct {
	// Cart
	// Discount
	// Genre
	// Ownership
	// Review
	// Role
	UserServiceInterface
	AuthServiceInterface
	// Wallet
	CompanyServiceInterface
	GameServiceInterface
}

func New(repo *repository.Repository) *Service {
	return &Service{
		UserServiceInterface:    NewUserService(repo.UserRepo),
		AuthServiceInterface:    NewAuthService(repo.AuthRepo),
		CompanyServiceInterface: NewCompanyService(repo.CompanyRepo),
		GameServiceInterface:    NewGameService(repo.GameRepo),
	}
}
