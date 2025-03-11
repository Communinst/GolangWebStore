package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type GenreService struct {
	repo repository.GenreRepo
}

func NewGenreService(repo repository.GenreRepo) *GenreService {
	return &GenreService{
		repo: repo,
	}
}

func (service *GenreService) AddGenre(ctx context.Context, name string, description string) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddGenre(c, name, description)
}

func (service *GenreService) GetGenreByName(ctx context.Context, name string) (*entities.Genre, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGenreByName(c, name)
}

func (service *GenreService) GetAllGenres(ctx context.Context) ([]entities.Genre, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetAllGenres(c)
}

func (service *GenreService) DeleteGenre(ctx context.Context, genreId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteGenre(c, genreId)
}

func (service *GenreService) GetGenreByID(ctx context.Context, genreId int) (*entities.Genre, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGenreByID(c, genreId)
}
