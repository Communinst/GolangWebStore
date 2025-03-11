package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type walletRepo struct {
	db *sqlx.DB
}

func NewWalletRepo(db *sqlx.DB) *walletRepo {
	return &walletRepo{
		db: db,
	}
}

func (repo *walletRepo) GetWalletByUserID(ctx context.Context, userId int) (*entities.Wallet, error) {
	var wallet entities.Wallet

	query := `SELECT * FROM wallets WHERE user_id = $1`

	err := repo.db.GetContext(ctx, &wallet, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &customErrors.ErrorWithStatusCode{
				HTTPStatus: http.StatusNotFound,
				Msg:        "Wallet not found",
			}
		}
		slog.Error("error retrieving wallet by user ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve wallet",
		}
	}

	log.Printf("Wallet retrieved for user ID %d", userId)
	return &wallet, nil
}

func (repo *walletRepo) UpdateWalletBalance(ctx context.Context, userId int, amount int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `UPDATE wallets SET balance = balance + $1 WHERE user_id = $2`

	_, err = tx.ExecContext(ctx, query, amount, userId)
	if err != nil {
		tx.Rollback()
		slog.Error("error updating wallet balance")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to update wallet balance",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Wallet balance updated for user ID %d", userId)
	return nil
}
