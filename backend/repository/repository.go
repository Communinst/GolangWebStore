package repository

import (
	"context"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable         = "users"
	gamesTable         = "games"
	companiesTable     = "companies"
	rolesTable         = "roles"
	storesTable        = "stores"
	categoriesTable    = "categories"
	manufacturersTable = "manufacturers"
	instrumentsTable   = "instruments"
	rentalsTable       = "rentals"
	paymentsTable      = "payments"
	repairsTable       = "repairs"
	discountsTable     = "repairs"
	reviewsTable       = "reviews"
)

// type AuthorizationRepo interface {
// }
// type CartRepo interface {
// }

// type DiscountRepo interface {
// }

// type GenreRepo interface {
// }
// type OwnershipRepo interface {
// }
// type ReviewRepo interface {
// }
// type RoleRepo interface {
// }
type UserRepo interface {
	PostUser(ctx context.Context, user *entities.User) (int, error)
	GetUser(ctx context.Context, userId int) (*entities.User, error)
	GetAllUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, userId int) error
	PutUserRole(ctx context.Context, userId int, roleId int) error
}

type GameRepo interface {
	PostGame(ctx context.Context, game *entities.Game) (int, error)
	GetGame(ctx context.Context, gameId int) (*entities.Game, error)
	GetAllGames(ctx context.Context) ([]entities.Game, error)
	DeleteGame(ctx context.Context, gameId int) error
	PutGamePrice(ctx context.Context, gameId int, price int) error
}

type CompanyRepo interface {
	PostCompany(ctx context.Context, company *entities.Company) (int, error)
	GetCompany(ctx context.Context, companyId int) (*entities.Company, error)
	GetAllCompanies(ctx context.Context) ([]entities.Company, error)
	DeleteCompany(ctx context.Context, companyId int) error
}

// type WalletRepo interface {
// }

type Repository struct {
	// AuthorizationRepo
	// CartRepo
	CompanyRepo
	// DiscountRepo
	GameRepo
	// GenreRepo
	// OwnershipRepo
	// ReviewRepo
	// RoleRepo
	UserRepo
	// WalletRepo
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:    NewUserRepo(db),
		GameRepo:    NewGameRepo(db),
		CompanyRepo: NewCompanyRepo(db),
	}
}
