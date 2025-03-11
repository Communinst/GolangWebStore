package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type DiscountService struct {
	repo repository.DiscountRepo
}

func NewDiscountService(repo repository.DiscountRepo) *DiscountService {
	return &DiscountService{
		repo: repo,
	}
}

func (service *DiscountService) AddDiscount(ctx context.Context, gameId int, discountValue int, startDate time.Time, ceaseDate time.Time) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddDiscount(c, gameId, discountValue, startDate, ceaseDate)
}

func (service *DiscountService) GetDiscountsByGameID(ctx context.Context, gameId int) ([]entities.Discount, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetDiscountsByGameID(c, gameId)
}

func (service *DiscountService) DeleteDiscount(ctx context.Context, discountId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteDiscount(c, discountId)
}
