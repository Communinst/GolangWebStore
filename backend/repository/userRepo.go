package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
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
		slog.Error("transaction initiation error")
		return -1, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`INSERT INTO %s (login, password, nickname, email, sing_up_date)
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
		slog.Error("transaction fulfillment error")
		return -1, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
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
		return &resultUser, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("user with %d if wasn't found", userId),
		}
	}

	slog.Error("unknown error obtaining user by it")
	return &resultUser, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (u *userRepo) getAllUsers(ctx context.Context) ([]entities.User, error) {
	var resultUsers []entities.User

	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)

	err := u.db.SelectContext(ctx, &resultUsers, query)
	if err == nil {
		return resultUsers, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return resultUsers, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "no user found",
		}
	}

	slog.Error("unknown error obtaining user by it")
	return resultUsers, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (u *userRepo) deleteUser(ctx context.Context, userId int) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1`, usersTable)

	result, err := tx.ExecContext(ctx, query, userId)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error deleting user by %d id", userId), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete user",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete user",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	if affectedAmount == 0 {
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("User by %d id wasn't found", userId),
		}
	}

	return nil
}

func (u *userRepo) putUserRole(ctx context.Context, userId int, roleId int) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`UPDATE %s
		SET role_id = $1
		WHERE user_id = $2`, usersTable)

	result, err := tx.ExecContext(ctx, query, roleId, userId)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error updating role of user by %d id", userId), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to update user's role",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to update user's role",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	if affectedAmount == 0 {
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("User by %d id wasn't found", userId),
		}
	}

	return nil
}
