package repositories

import (
	"context"

	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
)


func (pg *TravelRequestDB) Create(ctx context.Context, travelRequest *models.TravelRequest) error {
	err := pg.Session.WithContext(ctx).Create(&travelRequest).Error
	return err
}

