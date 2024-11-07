package repositories

import (
	"context"

	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)


func (pg *TravelRequestDB) Find(ctx context.Context, filter map[string]interface{}, pagination utils.Pagination) ([]*models.TravelRequest, error){
	var travelsRequest []*models.TravelRequest

	query := pg.Session.WithContext(ctx)

	if pagination.Page > 0 && pagination.Limit > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	if len(filter) > 0 {
		for key, value := range filter {
			if value != nil && value != "" {
				query = query.Where(key, value)
			}
		}
	}

	if err := query.Find(&travelsRequest).Error; err != nil {
		return nil, err
	}

	return travelsRequest, nil
	
}

