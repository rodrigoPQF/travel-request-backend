package repositories

import (
	"context"

	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
)

func (pg *TravelRequestDB) Update(ctx context.Context, travelRequest *models.TravelRequest) error {
	return pg.Session.WithContext(ctx).Save(&travelRequest).Error
}