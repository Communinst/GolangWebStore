// Describe overall entities and functions
// Get back to HANDLER_MAIN to determine the relevant routes for each endpoints
// If not done inside 24 hours, you're better be dead
// Freaking idiot

// Up handler

package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type Genre interface {
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
type WalletServiceInterface interface {
	GetWalletByUserID(ctx context.Context, userId int) (*entities.Wallet, error)
	UpdateWalletBalance(ctx context.Context, userId int, amount int64) error
}

type CompanyServiceInterface interface {
	PostCompany(ctx context.Context, company *entities.Company) error
	GetCompany(ctx context.Context, companyId int) (*entities.Company, error)
	GetAllCompanies(ctx context.Context) ([]entities.Company, error)
	DeleteCompany(ctx context.Context, companyId int) error
	GetCompanyByName(ctx context.Context, name string) (*entities.Company, error)
	DeleteCompanyByName(ctx context.Context, name string) error
}
type GameServiceInterface interface {
	PostGame(ctx context.Context, game *entities.Game) error
	GetGame(ctx context.Context, gameId int) (*entities.Game, error)
	GetGameByName(ctx context.Context, gameName string) (*entities.Game, error) // New method
	GetAllGames(ctx context.Context) ([]entities.Game, error)
	DeleteGame(ctx context.Context, gameId int) error
	DeleteGameByName(ctx context.Context, gameName string) error // New method
	PutGamePrice(ctx context.Context, gameId int, newPrice int) error
}
type GenreServiceInterface interface {
	AddGenre(ctx context.Context, name string, description string) error
	GetGenreByName(ctx context.Context, name string) (*entities.Genre, error)
	GetGenreByID(ctx context.Context, genreId int) (*entities.Genre, error)
	GetAllGenres(ctx context.Context) ([]entities.Genre, error)
	DeleteGenre(ctx context.Context, genreId int) error
}

type CartServiceInterface interface {
	AddGameToCart(ctx context.Context, userId int, gameId int) error
	GetCartByUserID(ctx context.Context, userId int) ([]entities.Game, error)
	RemoveGameFromCart(ctx context.Context, userId int, gameId int) error
}

type OwnershipServiceInterface interface {
	AddOwnership(ctx context.Context, userId int, gameId int, minutesSpent int64, receiptDate time.Time) error
	GetOwnershipsByUserID(ctx context.Context, userId int) ([]entities.Ownership, error)
	GetOwnershipsByGameID(ctx context.Context, gameId int) ([]entities.Ownership, error)
	DeleteOwnership(ctx context.Context, ownershipId int) error
}
type DiscountServiceInterface interface {
	AddDiscount(ctx context.Context, gameId int, discountValue int, startDate time.Time, ceaseDate time.Time) error
	GetDiscountsByGameID(ctx context.Context, gameId int) ([]entities.Discount, error)
	DeleteDiscount(ctx context.Context, discountId int) error
}

type ReviewServiceInterface interface {
	AddReview(ctx context.Context, userId int, gameId int, recommended bool, message string, date time.Time) error
	GetReviewsByGameID(ctx context.Context, gameId int) ([]entities.Review, error)
	GetReviewsByUserID(ctx context.Context, userId int) ([]entities.Review, error)
	DeleteReview(ctx context.Context, reviewId int) error
}

type Service struct {
	UserServiceInterface
	AuthServiceInterface
	WalletServiceInterface

	CompanyServiceInterface
	GameServiceInterface
	GenreServiceInterface

	CartServiceInterface
	OwnershipServiceInterface
	DiscountServiceInterface

	ReviewServiceInterface
}

func New(repo *repository.Repository) *Service {
	return &Service{
		UserServiceInterface:      NewUserService(repo.UserRepo),
		AuthServiceInterface:      NewAuthService(repo.AuthRepo),
		WalletServiceInterface:    NewWalletService(repo.WalletRepo),
		CompanyServiceInterface:   NewCompanyService(repo.CompanyRepo),
		GameServiceInterface:      NewGameService(repo.GameRepo),
		GenreServiceInterface:     NewGenreService(repo.GenreRepo),
		CartServiceInterface:      NewCartService(repo.CartRepo),
		OwnershipServiceInterface: NewOwnershipService(repo.OwnershipRepo),
		DiscountServiceInterface:  NewDiscountService(repo.DiscountRepo),
		ReviewServiceInterface:    NewReviewService(repo.ReviewRepo),
	}
}
