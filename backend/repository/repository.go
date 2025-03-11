package repository

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/jmoiron/sqlx"
)

const (
	rolesTable   = "roles"
	usersTable   = "users"
	walletsTable = "wallets"

	companiesTable   = "companies"
	genresTable      = "genres"
	gamesTable       = "games"
	gamesGenresTable = "games_genres"

	cartsTable     = "carts"
	cartGamesTable = "cart_games"
	ownershipTable = "ownerships"

	discountsTable = "discountst"
	reviewsTable   = "reviews"
)

// type CartRepo interface {
// }

// type GenreRepo interface {
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
type AuthRepo interface {
	PostUser(ctx context.Context, user *entities.User) (int, error)
	GetUser(ctx context.Context, userId int) (*entities.User, error)
	GetUserByEmail(ctx context.Context, userEmail string) (*entities.User, error)
}
type WalletRepo interface {
	GetWalletByUserID(ctx context.Context, userId int) (*entities.Wallet, error)
	UpdateWalletBalance(ctx context.Context, userId int, amount int64) error
}

type CompanyRepo interface {
	PostCompany(ctx context.Context, company *entities.Company) (int, error)
	GetCompany(ctx context.Context, companyId int) (*entities.Company, error)
	GetAllCompanies(ctx context.Context) ([]entities.Company, error)
	DeleteCompany(ctx context.Context, companyId int) error
	GetCompanyByName(ctx context.Context, name string) (*entities.Company, error)
	DeleteCompanyByName(ctx context.Context, name string) error
}
type GameRepo interface {
	PostGame(ctx context.Context, game *entities.Game) (int, error)
	GetGame(ctx context.Context, gameId int) (*entities.Game, error)
	GetGameByName(ctx context.Context, gameName string) (*entities.Game, error) // New method
	GetAllGames(ctx context.Context) ([]entities.Game, error)
	DeleteGame(ctx context.Context, gameId int) error
	DeleteGameByName(ctx context.Context, gameName string) error // New method
	PutGamePrice(ctx context.Context, gameId int, price int) error
}
type GenreRepo interface {
	AddGenre(ctx context.Context, name string, description string) error
	GetGenreByName(ctx context.Context, name string) (*entities.Genre, error)
	GetGenreByID(ctx context.Context, genreId int) (*entities.Genre, error)
	GetAllGenres(ctx context.Context) ([]entities.Genre, error)
	DeleteGenre(ctx context.Context, genreId int) error
}
type GameGenreRepo interface {
	AddGenreToGame(ctx context.Context, gameId int, genreId int) error
	GetGenresByGameID(ctx context.Context, gameId int) ([]entities.Genre, error)
	GetGamesByGenreID(ctx context.Context, genreId int) ([]entities.Game, error)
	GetGamesByGenreName(ctx context.Context, genreName string) ([]entities.Game, error)
	IncrementGenreCount(ctx context.Context, gameId int, genreId int) error
	DeleteGameGenre(ctx context.Context, gameId int, genreId int) error
}

type CartRepo interface {
	AddGameToCart(ctx context.Context, userId int, gameId int) error
	GetCartByUserID(ctx context.Context, userId int) ([]entities.Game, error)
	RemoveGameFromCart(ctx context.Context, userId int, gameId int) error
}
type OwnershipRepo interface {
	AddOwnership(ctx context.Context, userId int, gameId int, minutesSpent int64, receiptDate time.Time) error
	GetOwnershipsByUserID(ctx context.Context, userId int) ([]entities.Ownership, error)
	GetOwnershipsByGameID(ctx context.Context, gameId int) ([]entities.Ownership, error)
	DeleteOwnership(ctx context.Context, ownershipId int) error
}
type DiscountRepo interface {
	AddDiscount(ctx context.Context, gameId int, discountValue int, startDate time.Time, ceaseDate time.Time) error
	GetDiscountsByGameID(ctx context.Context, gameId int) ([]entities.Discount, error)
	DeleteDiscount(ctx context.Context, discountId int) error
}

type ReviewRepo interface {
	AddReview(ctx context.Context, userId int, gameId int, recommended bool, message string, date time.Time) error
	GetReviewsByGameID(ctx context.Context, gameId int) ([]entities.Review, error)
	GetReviewsByUserID(ctx context.Context, userId int) ([]entities.Review, error)
	DeleteReview(ctx context.Context, reviewId int) error
}

type DumpRepo interface {
	InsertDump(ctx context.Context, dump *entities.Dump) error
	GetAllDumps(ctx context.Context) ([]entities.Dump, error)
}

type Repository struct {
	// RoleRepo
	UserRepo
	AuthRepo
	WalletRepo
	CompanyRepo
	GameRepo
	GenreRepo
	GameGenreRepo
	CartRepo
	OwnershipRepo
	DiscountRepo
	ReviewRepo
	DumpRepo
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:      NewUserRepo(db),
		AuthRepo:      NewAuthRepo(db),
		WalletRepo:    NewWalletRepo(db),
		GameRepo:      NewGameRepo(db),
		GenreRepo:     NewGenreRepo(db),
		GameGenreRepo: NewGameGenreRepo(db),
		CompanyRepo:   NewCompanyRepo(db),
		CartRepo:      NewCartRepo(db),
		OwnershipRepo: NewOwnershipRepo(db),
		DiscountRepo:  NewDiscountRepo(db),
		ReviewRepo:    NewReviewRepo(db),
		DumpRepo:      NewDumpRepo(db),
	}
}
