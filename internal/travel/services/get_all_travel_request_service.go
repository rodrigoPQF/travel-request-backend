package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

func (trs *travelRequestService) GetAll(ctx context.Context, input *dtos.GetAllTravelRequestInputDto) ([]*dtos.TravelRequestDto, error){
	result := make(map[string]interface{})

	result["status"] = input.Status

	res, err := trs.TravelRequestRepository.Find(ctx, result, utils.Pagination{
		Page: input.Page,
		Limit: input.Limit,
	})
	
	if err != nil {
		return nil, err
	}

	var out []*dtos.TravelRequestDto
	copier.Copy(&out, res)
	
	return out, nil

}