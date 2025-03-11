package repository

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type ownershipRepo struct {
	db *sqlx.DB
}

func NewOwnershipRepo(db *sqlx.DB) *ownershipRepo {
	return &ownershipRepo{
		db: db,
	}
}

func (repo *ownershipRepo) AddOwnership(ctx context.Context, userId int, gameId int, minutesSpent int64, receiptDate time.Time) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO ownerships (user_id, game_id, minutes_spent, receipt_date) VALUES ($1, $2, $3, $4)`

	_, err = tx.ExecContext(ctx, query, userId, gameId, minutesSpent, receiptDate)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding ownership")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add ownership",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Ownership added for user ID %d and game ID %d", userId, gameId)
	return nil
}

func (repo *ownershipRepo) GetOwnershipsByUserID(ctx context.Context, userId int) ([]entities.Ownership, error) {
	var ownerships []entities.Ownership

	query := `SELECT * FROM ownerships WHERE user_id = $1`

	err := repo.db.SelectContext(ctx, &ownerships, query, userId)
	if err != nil {
		slog.Error("error retrieving ownerships by user ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve ownerships",
		}
	}

	log.Printf("Ownerships retrieved for user ID %d", userId)
	return ownerships, nil
}

func (repo *ownershipRepo) GetOwnershipsByGameID(ctx context.Context, gameId int) ([]entities.Ownership, error) {
	var ownerships []entities.Ownership

	query := `SELECT * FROM ownerships WHERE game_id = $1`

	err := repo.db.SelectContext(ctx, &ownerships, query, gameId)
	if err != nil {
		slog.Error("error retrieving ownerships by game ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve ownerships",
		}
	}

	log.Printf("Ownerships retrieved for game ID %d", gameId)
	return ownerships, nil
}

func (repo *ownershipRepo) DeleteOwnership(ctx context.Context, ownershipId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `DELETE FROM ownerships WHERE ownership_id = $1`

	result, err := tx.ExecContext(ctx, query, ownershipId)
	if err != nil {
		tx.Rollback()
		slog.Error("error deleting ownership")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete ownership",
		}
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete ownership",
		}
	}

	if affected == 0 {
		tx.Rollback()
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        "Ownership not found",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction commit error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction commit failed",
		}
	}

	log.Printf("Ownership with ID %d deleted successfully", ownershipId)
	return nil
}
