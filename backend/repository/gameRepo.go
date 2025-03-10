package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type gameRepo struct {
	db *sqlx.DB
}

func NewGameRepo(db *sqlx.DB) *gameRepo {
	return &gameRepo{
		db: db,
	}
}

func (repo *gameRepo) PostGame(ctx context.Context, game *entities.Game) (int, error) {
	var resultId int
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return -1, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`INSERT INTO %s (publisher_id, name, description, price, release_date, rating)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING game_id`, gamesTable)

	err = tx.QueryRowContext(ctx, query,
		game.PublisherId,
		game.Name,
		game.Description,
		game.Price,
		game.Releasedate,
		game.Rating).Scan(&resultId)

	if err != nil {
		log.Printf("%d, %s", game.PublisherId, err.Error())
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

	log.Printf("Game with name %s posted successfully", game.Name)
	return resultId, err
}

func (repo *gameRepo) GetGame(ctx context.Context, gameId int) (*entities.Game, error) {
	var resultGame entities.Game

	query := fmt.Sprintf(`SELECT * FROM %s WHERE game_id = $1`, gamesTable)

	err := repo.db.GetContext(ctx, &resultGame, query, gameId)
	if err == nil {
		log.Printf("Game with id %d was obtained", gameId)
		return &resultGame, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("game with id %d was not found", gameId),
		}
	}

	slog.Error("unknown error obtaining user by id", "error", err)
	return nil, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown internal server error occurred",
	}
}

func (repo *gameRepo) GetAllGames(ctx context.Context) ([]entities.Game, error) {
	var resultGames []entities.Game

	query := fmt.Sprintf(`SELECT * FROM %s`, gamesTable)

	err := repo.db.SelectContext(ctx, &resultGames, query)
	if err == nil {
		return resultGames, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return resultGames, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "no game found",
		}
	}

	slog.Error("unknown error obtaining game by id")
	return resultGames, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (repo *gameRepo) DeleteGame(ctx context.Context, gameId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE game_id = $1`, gamesTable)

	result, err := tx.ExecContext(ctx, query, gameId)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error deleting game by %d id", gameId), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game",
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
			Msg:        fmt.Sprintf("User by %d id wasn't found", gameId),
		}
	}

	log.Printf("Game by %d id was deleted", gameId)
	return nil
}

func (repo *gameRepo) PutGamePrice(ctx context.Context, gameId int, price int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`UPDATE %s
		SET price = $1
		WHERE game_id = $2`, gamesTable)

	result, err := tx.ExecContext(ctx, query, price, gameId)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error updating price of game with %d id", gameId), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to update game's price",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to update game's price",
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
			Msg:        fmt.Sprintf("Game by %d id wasn't found", gameId),
		}
	}

	log.Printf("Game's by %d id price updated", gameId)
	return nil
}

func (repo *gameRepo) GetGameByName(ctx context.Context, gameName string) (*entities.Game, error) {
	var resultGame entities.Game

	query := fmt.Sprintf(`SELECT * FROM %s WHERE name = $1`, gamesTable)

	err := repo.db.GetContext(ctx, &resultGame, query, gameName)
	if err == nil {
		log.Printf("Game by name %s was obtained", gameName)
		return &resultGame, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        fmt.Sprintf("game with name %s wasn't found", gameName),
		}
	}

	slog.Error("unknown error obtaining game by name")
	return nil, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown internal server error occurred",
	}
}

func (repo *gameRepo) DeleteGameByName(ctx context.Context, gameName string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE name = $1`, gamesTable)

	result, err := tx.ExecContext(ctx, query, gameName)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error deleting game by name: %s", gameName), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game",
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
			Msg:        fmt.Sprintf("Game with name %s wasn't found", gameName),
		}
	}

	log.Printf("Game with name %s was deleted", gameName)
	return nil
}
