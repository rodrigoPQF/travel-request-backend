package repositories

import (
	"context"

	"github.com/rodrigoPQF/travel-request-backend/internal/infra/database"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

type TravelRequestDB struct {
	*database.PostgresDB
}


type TravelRequestRepository interface {
	Tx(database.RepositoryTransaction) error

	Create(ctx context.Context, travelRequest *models.TravelRequest) error
	Update(ctx context.Context, travelRequest *models.TravelRequest) error
	FindById(ctx context.Context, id string) (*models.TravelRequest, error)
	Find(ctx context.Context, filter map[string]interface{}, pagination utils.Pagination) ([]*models.TravelRequest, error)
}