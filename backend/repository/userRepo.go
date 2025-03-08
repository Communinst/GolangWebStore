package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func newUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) postUser(ctx context.Context, user *entities.User) (int, error) {
	var result_id int

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	query := fmt.Sprintf(`INSERT INTO users (login, password, nickname, email, sing_up_date)
		VALUES ($1 $2 $3 $4 $5) RETURNING user_id`, usersTable)

	err = tx.QueryRowContext(ctx, query,
		user.Login,
		user.Password,
		user.Nickname,
		user.Email,
		user.SingUpDate).Scan(&result_id)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return result_id, err
}

func (u *userRepo) getUser(ctx context.Context, userId int) (*entities.User, error) {
	var resultUser entities.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1`, usersTable)

	err := u.db.GetContext(ctx, &resultUser, query, userId)
	if err == nil {
		return &resultUser, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return &resultUser, httpErrors
	}
}
