package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/http/handlers"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

// UpdateTravelRequest godoc
// @Summary Updates a travel request
// @Description Updates details of a travel request by ID
// @Tags Travel Request
// @Accept json
// @Produce json
// @Param id path string true "Travel Request ID"
// @Param input body dtos.UpdateTravelRequestBodyInputDto true "Request body for updating a travel request"
// @Success 200 {object} handlers.BaseResponseHandler
// @Failure 400 {object} handlers.BaseResponseHandler
// @Router /request [put]
func (trh *travelRequestHandler) UpdateTravelRequest(ctx *fiber.Ctx) error {

	travelId := ctx.Params("id")

	if travelId == "" {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, "an error occurred, id parameter is required")
	}

	input := dtos.UpdateTravelRequestInputDto{
		UpdateTravelRequestParamsInputDto: dtos.UpdateTravelRequestParamsInputDto{Id: travelId},
	}

	if err := ctx.BodyParser(&input.UpdateTravelRequestBodyInputDto); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if err := utils.ValidateStruct(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err := trh.travelRequestService.Update(ctx.Context(), &input)

	if err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx,http.StatusOK, "success update travel request", nil)
}