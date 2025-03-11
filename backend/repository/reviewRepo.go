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

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepo(db *sqlx.DB) *reviewRepo {
	return &reviewRepo{
		db: db,
	}
}

func (repo *reviewRepo) AddReview(ctx context.Context, userId int, gameId int, recommended bool, message string, date time.Time) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO reviews (user_id, game_id, recommended, message, date) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.ExecContext(ctx, query, userId, gameId, recommended, message, date)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding review")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add review",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Review added for user ID %d and game ID %d", userId, gameId)
	return nil
}

func (repo *reviewRepo) GetReviewsByGameID(ctx context.Context, gameId int) ([]entities.Review, error) {
	var reviews []entities.Review

	query := `SELECT * FROM reviews WHERE game_id = $1`

	err := repo.db.SelectContext(ctx, &reviews, query, gameId)
	if err != nil {
		slog.Error("error retrieving reviews by game ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve reviews",
		}
	}

	log.Printf("Reviews retrieved for game ID %d", gameId)
	return reviews, nil
}

func (repo *reviewRepo) GetReviewsByUserID(ctx context.Context, userId int) ([]entities.Review, error) {
	var reviews []entities.Review

	query := `SELECT * FROM reviews WHERE user_id = $1`

	err := repo.db.SelectContext(ctx, &reviews, query, userId)
	if err != nil {
		slog.Error("error retrieving reviews by user ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve reviews",
		}
	}

	log.Printf("Reviews retrieved for user ID %d", userId)
	return reviews, nil
}

func (repo *reviewRepo) DeleteReview(ctx context.Context, reviewId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `DELETE FROM reviews WHERE review_id = $1`

	result, err := tx.ExecContext(ctx, query, reviewId)
	if err != nil {
		tx.Rollback()
		slog.Error("error deleting review")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete review",
		}
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete review",
		}
	}

	if affected == 0 {
		tx.Rollback()
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        "Review not found",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction commit error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction commit failed",
		}
	}

	log.Printf("Review with ID %d deleted successfully", reviewId)
	return nil
}
