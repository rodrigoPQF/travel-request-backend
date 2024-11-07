package services

import (
	"context"

	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/repositories"
)

type TravelRequestService interface {
	Create(ctx context.Context, input *dtos.CreateTravelRequestInputDto) (*dtos.CreateTravelRequestOutputDto, error)
	Update(ctx context.Context, input *dtos.UpdateTravelRequestInputDto) error
	Get(ctx context.Context, input *dtos.GetTravelRequestInputDto) (*dtos.TravelRequestDto, error)
	GetAll(ctx context.Context, input *dtos.GetAllTravelRequestInputDto) ([]*dtos.TravelRequestDto, error)
}

type travelRequestService struct {
	TravelRequestRepository repositories.TravelRequestRepository
}

func NewTravelRequestService(repository repositories.TravelRequestRepository) TravelRequestService {
	return &travelRequestService{
		TravelRequestRepository: repository,
	}
}
