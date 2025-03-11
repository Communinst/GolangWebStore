package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type GameGenreService struct {
	repo repository.GameGenreRepo
}

func NewGameGenreService(repo repository.GameGenreRepo) *GameGenreService {
	return &GameGenreService{
		repo: repo,
	}
}

func (service *GameGenreService) AddGenreToGame(ctx context.Context, gameId int, genreId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddGenreToGame(c, gameId, genreId)
}

func (service *GameGenreService) GetGenresByGameID(ctx context.Context, gameId int) ([]entities.Genre, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGenresByGameID(c, gameId)
}

func (service *GameGenreService) GetGamesByGenreID(ctx context.Context, genreId int) ([]entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGamesByGenreID(c, genreId)
}

func (service *GameGenreService) GetGamesByGenreName(ctx context.Context, genreName string) ([]entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGamesByGenreName(c, genreName)
}

func (service *GameGenreService) IncrementGenreCount(ctx context.Context, gameId int, genreId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.IncrementGenreCount(c, gameId, genreId)
}

func (service *GameGenreService) DeleteGameGenre(ctx context.Context, gameId int, genreId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteGameGenre(c, gameId, genreId)
}
