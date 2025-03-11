package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type WalletService struct {
	repo repository.WalletRepo
}

func NewWalletService(repo repository.WalletRepo) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (service *WalletService) GetWalletByUserID(ctx context.Context, userId int) (*entities.Wallet, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetWalletByUserID(c, userId)
}

func (service *WalletService) UpdateWalletBalance(ctx context.Context, userId int, amount int64) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.UpdateWalletBalance(c, userId, amount)
}
