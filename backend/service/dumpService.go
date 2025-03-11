package service

import (
	"context"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
)

type DumpService struct {
	repo repository.DumpRepo
}

func NewDumpService(repo repository.DumpRepo) *DumpService {
	return &DumpService{repo: repo}
}

func (s *DumpService) InsertDump(ctx context.Context, filePath string) error {
	dump := &entities.Dump{
		Filename:  filePath,
		CreatedAt: time.Now(),
	}
	return s.repo.InsertDump(ctx, dump)
}

func (s *DumpService) GetAllDumps(ctx context.Context) ([]entities.Dump, error) {
	return s.repo.GetAllDumps(ctx)
}
