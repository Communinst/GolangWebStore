package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type GameService struct {
	repo repository.GameRepo
}

func NewGameService(repo repository.GameRepo) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (service *GameService) PostGame(ctx context.Context, game *entities.Game) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	_, err := service.repo.PostGame(c, game)
	return err
}

func (service *GameService) GetGame(ctx context.Context, gameId int) (*entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGame(c, gameId)
}

func (service *GameService) GetAllGames(ctx context.Context) ([]entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetAllGames(c)
}

func (service *GameService) DeleteGame(ctx context.Context, gameId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteGame(c, gameId)
}

func (service *GameService) PutGamePrice(ctx context.Context, gameId int, newPrice int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.PutGamePrice(c, gameId, newPrice)
}
func (service *GameService) GetGameByName(ctx context.Context, gameName string) (*entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetGameByName(c, gameName)
}

func (service *GameService) DeleteGameByName(ctx context.Context, gameName string) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteGameByName(c, gameName)
}
