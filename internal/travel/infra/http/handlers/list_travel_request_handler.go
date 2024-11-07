package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/http/handlers"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

// ListTravelRequest godoc
// @Summary Retrieves a travel request
// @Description Fetches details of a travel request by ID
// @Tags Travel Request
// @Accept json
// @Produce json
// @Param id path string true "Travel Request ID"
// @Success 200 {object} handlers.BaseResponseHandler{payload=dtos.TravelRequestDto}
// @Failure 400 {object} handlers.BaseResponseHandler
// @Router /request/{id} [get]
func (trh *travelRequestHandler) ListTravelRequest(ctx *fiber.Ctx) error {

	travelId := ctx.Params("id")

	if travelId == "" {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, "an error occurred, id parameter is required")
	}

	input := dtos.GetTravelRequestInputDto{
		Id: travelId,
	}

	if err := utils.ValidateStruct(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := trh.travelRequestService.Get(ctx.Context(), &input)
	
	if err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx,http.StatusOK, "success find travel request", output)
}