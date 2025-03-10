package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type CompanyService struct {
	repo repository.CompanyRepo
}

func NewCompanyService(repo repository.CompanyRepo) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (service *CompanyService) PostCompany(ctx context.Context, company *entities.Company) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	_, err := service.repo.PostCompany(c, company)
	return err
}

func (service *CompanyService) GetCompany(ctx context.Context, companyId int) (*entities.Company, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetCompany(c, companyId)
}

func (service *CompanyService) GetAllCompanies(ctx context.Context) ([]entities.Company, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.GetAllCompanies(c)
}

func (service *CompanyService) DeleteCompany(ctx context.Context, companyId int) error {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return service.repo.DeleteCompany(c, companyId)
}
