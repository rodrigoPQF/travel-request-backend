package services

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
)


func (trs *travelRequestService) Update(ctx context.Context, input *dtos.UpdateTravelRequestInputDto) error {

  travelRequestExist, err := trs.TravelRequestRepository.FindById(ctx, input.Id)

	if err != nil {
		return err
	}
	if travelRequestExist == nil {
		return errors.New("error travel request not exist")
	}

	if err := copier.Copy(&travelRequestExist, &input); err != nil {
		return err
	}

	return trs.TravelRequestRepository.Update(ctx, travelRequestExist)
}