package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type ReviewService struct {
	repo repository.ReviewRepo
}

func NewReviewService(repo repository.ReviewRepo) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (service *ReviewService) AddReview(ctx context.Context, userId int, gameId int, recommended bool, message string, date time.Time) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.AddReview(c, userId, gameId, recommended, message, date)
}

func (service *ReviewService) GetReviewsByGameID(ctx context.Context, gameId int) ([]entities.Review, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetReviewsByGameID(c, gameId)
}

func (service *ReviewService) GetReviewsByUserID(ctx context.Context, userId int) ([]entities.Review, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetReviewsByUserID(c, userId)
}

func (service *ReviewService) DeleteReview(ctx context.Context, reviewId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteReview(c, reviewId)
}
