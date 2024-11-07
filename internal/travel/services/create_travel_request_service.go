package services

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
)


func (trs *travelRequestService) Create(ctx context.Context, input *dtos.CreateTravelRequestInputDto) (*dtos.CreateTravelRequestOutputDto, error) {
	travelRequestExist, err := trs.TravelRequestRepository.FindById(ctx, input.Id)

	if err != nil {
		return nil, err
	}

	if travelRequestExist != nil {
		return nil, errors.New("error travel request already exist")
	}
	
	var travelRequest models.TravelRequest

	copier.Copy(&travelRequest, &input)

	travelRequest.SetDates(input.DepartureDate, input.ReturnDate)

	if err := trs.TravelRequestRepository.Create(ctx, &travelRequest); err != nil {
		return nil, err
	}

	return &dtos.CreateTravelRequestOutputDto{}, nil
}
