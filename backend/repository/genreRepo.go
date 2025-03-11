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

type genreRepo struct {
	db *sqlx.DB
}

func NewGenreRepo(db *sqlx.DB) *genreRepo {
	return &genreRepo{
		db: db,
	}
}

func (repo *genreRepo) AddGenre(ctx context.Context, name string, description string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `INSERT INTO genres (name, description) VALUES ($1, $2)`

	_, err = tx.ExecContext(ctx, query, name, description)
	if err != nil {
		tx.Rollback()
		slog.Error("error adding genre")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to add genre",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Genre added: %s", name)
	return nil
}

func (repo *genreRepo) GetGenreByName(ctx context.Context, name string) (*entities.Genre, error) {
	var genre entities.Genre

	query := `SELECT * FROM genres WHERE name = $1`

	err := repo.db.GetContext(ctx, &genre, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &customErrors.ErrorWithStatusCode{
				HTTPStatus: http.StatusNotFound,
				Msg:        "Genre not found",
			}
		}
		slog.Error("error retrieving genre by name")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve genre",
		}
	}

	log.Printf("Genre retrieved: %s", name)
	return &genre, nil
}

func (repo *genreRepo) GetAllGenres(ctx context.Context) ([]entities.Genre, error) {
	var genres []entities.Genre
	log.Print("In here")
	query := fmt.Sprintf(`SELECT * FROM %s`, genresTable)

	err := repo.db.SelectContext(ctx, &genres, query)
	if err == nil {
		log.Print("Users were obtained")
		return genres, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return genres, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "no genre found",
		}
	}

	slog.Error("unknown error obtaining genres")
	return genres, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (repo *genreRepo) DeleteGenre(ctx context.Context, genreId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := `DELETE FROM genres WHERE genre_id = $1`

	result, err := tx.ExecContext(ctx, query, genreId)
	if err != nil {
		tx.Rollback()
		slog.Error("error deleting genre")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete genre",
		}
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting affected rows")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete genre",
		}
	}

	if affected == 0 {
		tx.Rollback()
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusNotFound,
			Msg:        "Genre not found",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction commit error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction commit failed",
		}
	}

	log.Printf("Genre with ID %d deleted successfully", genreId)
	return nil
}

func (repo *genreRepo) GetGenreByID(ctx context.Context, genreId int) (*entities.Genre, error) {
	var genre entities.Genre

	query := `SELECT * FROM genres WHERE genre_id = $1`

	err := repo.db.GetContext(ctx, &genre, query, genreId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &customErrors.ErrorWithStatusCode{
				HTTPStatus: http.StatusNotFound,
				Msg:        "Genre not found",
			}
		}
		slog.Error("error retrieving genre by ID")
		return nil, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to retrieve genre",
		}
	}

	log.Printf("Genre retrieved with ID: %d", genreId)
	return &genre, nil
}
