package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type OwnershipService struct {
	repo repository.OwnershipRepo
}

func NewOwnershipService(repo repository.OwnershipRepo) *OwnershipService {
	return &OwnershipService{
		repo: repo,
	}
}

func (service *OwnershipService) AddOwnership(ctx context.Context, userId int, gameId int, minutesSpent int64, receiptDate time.Time) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddOwnership(c, userId, gameId, minutesSpent, receiptDate)
}

func (service *OwnershipService) GetOwnershipsByUserID(ctx context.Context, userId int) ([]entities.Ownership, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetOwnershipsByUserID(c, userId)
}

func (service *OwnershipService) GetOwnershipsByGameID(ctx context.Context, gameId int) ([]entities.Ownership, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetOwnershipsByGameID(c, gameId)
}

func (service *OwnershipService) DeleteOwnership(ctx context.Context, ownershipId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteOwnership(c, ownershipId)
}
