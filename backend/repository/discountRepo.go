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

type discountRepo struct {
	db *sqlx.DB
}

func NewDiscountRepo(db *sqlx.DB) *discountRepo {
	return &discountRepo{
		db: db,
	}
}

func (repo *discountRepo) AddDiscount(ctx context.Context, gameId int, discountValue int, startDate time.Time, ceaseDate time.Time) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO discounts (game_id, discount_value, start_date, cease_date) VALUES ($1, $2, $3, $4)`

	_, err = tx.ExecContext(ctx, query, gameId, discountValue, startDate, ceaseDate)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding discount")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add discount",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Discount added for game ID %d", gameId)
	return nil
}

func (repo *discountRepo) GetDiscountsByGameID(ctx context.Context, gameId int) ([]entities.Discount, error) {
	var discounts []entities.Discount

	query := `SELECT * FROM discounts WHERE game_id = $1`

	err := repo.db.SelectContext(ctx, &discounts, query, gameId)
	if err != nil {
		slog.Error("error retrieving discounts by game ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve discounts",
		}
	}

	log.Printf("Discounts retrieved for game ID %d", gameId)
	return discounts, nil
}

func (repo *discountRepo) DeleteDiscount(ctx context.Context, discountId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `DELETE FROM discounts WHERE discount_id = $1`

	result, err := tx.ExecContext(ctx, query, discountId)
	if err != nil {
		tx.Rollback()
		slog.Error("error deleting discount")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete discount",
		}
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete discount",
		}
	}

	if affected == 0 {
		tx.Rollback()
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        "Discount not found",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction commit error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction commit failed",
		}
	}

	log.Printf("Discount with ID %d deleted successfully", discountId)
	return nil
}
