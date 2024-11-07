package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
	"gorm.io/gorm"
)

func (pg *TravelRequestDB) FindById(ctx context.Context, id string) (*models.TravelRequest, error){
	var travelRequest models.TravelRequest

	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	if err := pg.Session.WithContext(ctx).First(&travelRequest, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &travelRequest, nil
}
