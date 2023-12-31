package services

import (
	"context"

	"github.com/Moji00f/SimpleProject/api/dto"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/data/db"
	"github.com/Moji00f/SimpleProject/data/models"
	"github.com/Moji00f/SimpleProject/pkg/logging"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			preloads: []preload{{string: "Cities"}},
		},
	}
}

// Create
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}
