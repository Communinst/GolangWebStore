package repository

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

func newRepo() *Repository {
	return &Repository{}
}
