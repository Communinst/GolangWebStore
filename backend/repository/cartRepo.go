package repository

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type cartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) *cartRepo {
	return &cartRepo{
		db: db,
	}
}

func (repo *cartRepo) AddGameToCart(ctx context.Context, userId int, gameId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO cart_games (cart_id, game_id)
		VALUES ((SELECT cart_id FROM carts WHERE user_id = $1), $2)`

	_, err = tx.ExecContext(ctx, query, userId, gameId)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding game to cart")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add game to cart",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Game added to cart for user ID %d", userId)
	return nil
}

func (repo *cartRepo) GetCartByUserID(ctx context.Context, userId int) ([]entities.Game, error) {
	var games []entities.Game

	query := `SELECT g.* FROM games g
		JOIN cart_games cg ON g.game_id = cg.game_id
		JOIN carts c ON cg.cart_id = c.cart_id
		WHERE c.user_id = $1`

	err := repo.db.SelectContext(ctx, &games, query, userId)
	if err != nil {
		slog.Error("error retrieving cart games")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve cart games",
		}
	}

	log.Printf("Cart retrieved for user ID %d", userId)
	return games, nil
}

func (repo *cartRepo) RemoveGameFromCart(ctx context.Context, userId int, gameId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `DELETE FROM cart_games
		WHERE cart_id = (SELECT cart_id FROM carts WHERE user_id = $1)
		AND game_id = $2`

	result, err := tx.ExecContext(ctx, query, userId, gameId)
	if err != nil {
		tx.Rollback()
		slog.Error("error removing game from cart")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to remove game from cart",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to remove game from cart",
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
			HTTPStatus: http.StatusNotFound,
			Msg:        "Game not found in cart",
		}
	}

	log.Printf("Game removed from cart for user ID %d", userId)
	return nil
}
