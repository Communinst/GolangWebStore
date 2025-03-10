package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type CartService struct {
	repo repository.CartRepo
}

func NewCartService(repo repository.CartRepo) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (service *CartService) AddGameToCart(ctx context.Context, userId int, gameId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddGameToCart(c, userId, gameId)
}

func (service *CartService) GetCartByUserID(ctx context.Context, userId int) ([]entities.Game, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetCartByUserID(c, userId)
}

func (service *CartService) RemoveGameFromCart(ctx context.Context, userId int, gameId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.RemoveGameFromCart(c, userId, gameId)
}
