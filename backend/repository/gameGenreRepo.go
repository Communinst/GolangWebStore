package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type gameGenreRepo struct {
	db *sqlx.DB
}

func NewGameGenreRepo(db *sqlx.DB) *gameGenreRepo {
	return &gameGenreRepo{
		db: db,
	}
}

func (repo *gameGenreRepo) AddGenreToGame(ctx context.Context, gameId int, genreId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO games_genres (game_id, genre_id) VALUES ($1, $2) ON CONFLICT (game_id, genre_id) DO NOTHING`

	_, err = tx.ExecContext(ctx, query, gameId, genreId)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding genre to game")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add genre to game",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Genre added to game ID %d", gameId)
	return nil
}

func (repo *gameGenreRepo) GetGenresByGameID(ctx context.Context, gameId int) ([]entities.Genre, error) {
	var genres []entities.Genre

	query := `
		SELECT g.*
		FROM genres g
		JOIN games_genres gg ON g.genre_id = gg.genre_id
		WHERE gg.game_id = $1
	`

	err := repo.db.SelectContext(ctx, &genres, query, gameId)
	if err != nil {
		slog.Error("error retrieving genres by game ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve genres",
		}
	}

	log.Printf("Genres retrieved for game ID %d", gameId)
	return genres, nil
}

func (repo *gameGenreRepo) GetGamesByGenreID(ctx context.Context, genreId int) ([]entities.Game, error) {
	var games []entities.Game

	query := `
		SELECT g.*
		FROM games g
		JOIN games_genres gg ON g.game_id = gg.game_id
		WHERE gg.genre_id = $1
	`

	err := repo.db.SelectContext(ctx, &games, query, genreId)
	if err != nil {
		slog.Error("error retrieving games by genre ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve games",
		}
	}

	log.Printf("Games retrieved for genre ID %d", genreId)
	return games, nil
}

func (repo *gameGenreRepo) GetGamesByGenreName(ctx context.Context, genreName string) ([]entities.Game, error) {
	var games []entities.Game

	query := `
		SELECT g.*
		FROM games g
		JOIN games_genres gg ON g.game_id = gg.game_id
		JOIN genres ge ON gg.genre_id = ge.genre_id
		WHERE ge.name = $1
	`

	err := repo.db.SelectContext(ctx, &games, query, genreName)
	if err != nil {
		slog.Error("error retrieving games by genre name")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve games",
		}
	}

	log.Printf("Games retrieved for genre name %s", genreName)
	return games, nil
}

func (repo *gameGenreRepo) IncrementGenreCount(ctx context.Context, gameId int, genreId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `UPDATE games_genres SET count = count + 1 WHERE game_id = $1 AND genre_id = $2`

	_, err = tx.ExecContext(ctx, query, gameId, genreId)
	if err != nil {
		tx.Rollback()
		slog.Error("error incrementing genre count")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to increment genre count",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Genre count incremented for game ID %d and genre ID %d", gameId, genreId)
	return nil
}

func (repo *gameGenreRepo) DeleteGameGenre(ctx context.Context, gameId int, genreId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`UPDATE %s SET count = count - 1`, gamesGenresTable)

	result, err := tx.ExecContext(ctx, query, gameId, genreId)
	if err != nil {
		tx.Rollback()
		slog.Error("error deleting game genre")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game genre",
		}
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete game genre",
		}
	}

	if affected == 0 {
		tx.Rollback()
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        "Game genre not found",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction commit error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction commit failed",
		}
	}

	log.Printf("Game genre with game ID %d and genre ID %d deleted successfully", gameId, genreId)
	return nil
}
