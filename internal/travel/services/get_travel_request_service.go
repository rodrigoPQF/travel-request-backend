package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
)


func(trs *travelRequestService) Get(ctx context.Context, input *dtos.GetTravelRequestInputDto) (*dtos.TravelRequestDto, error){

	res, err := trs.TravelRequestRepository.FindById(ctx, input.Id)

	if err != nil {
		return nil, err
	}

	var out dtos.TravelRequestDto
	copier.Copy(&out, res)

	return &out, nil
}
