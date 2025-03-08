package repository

import (
	"context"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable         = "users"
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

type AuthorizationRepo interface {
}
type CartRepo interface {
}
type CompanyRepo interface {
}
type DiscountRepo interface {
}
type GameRepo interface {
}
type GenreRepo interface {
}
type OwnershipRepo interface {
}
type ReviewRepo interface {
}
type RoleRepo interface {
}
type UserRepo interface {
	postUser(ctx context.Context, user *entities.User) (int, error)
	getUser(ctx context.Context, userId int) (*entities.User, error)
	getAllUsers(ctx context.Context) ([]entities.User, error)
	deleteUser(ctx context.Context, userId int) error
	putUserRole(ctx context.Context, userId int, roleId int) error
}
type WalletRepo interface {
}

type Repository struct {
	AuthorizationRepo
	CartRepo
	CompanyRepo
	DiscountRepo
	GameRepo
	GenreRepo
	OwnershipRepo
	ReviewRepo
	RoleRepo
	UserRepo
	WalletRepo
}

func newRepsitory(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo: newUserRepo(db),
	}
}
