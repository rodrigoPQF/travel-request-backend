package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/http/handlers"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

// CreateTravelRequest godoc
// @Summary Creates a new travel request
// @Description Creates a new travel request with the provided details
// @Tags Travel Request
// @Accept json
// @Produce json
// @Param input body dtos.CreateTravelRequestInputDto true "Travel Request creation payload"
// @Success 201 {object} handlers.BaseResponseHandler
// @Failure 400 {object} handlers.BaseResponseHandler
// @Router /request [post]
func (trh *travelRequestHandler) CreateTravelRequest(ctx *fiber.Ctx) error {
	var input dtos.CreateTravelRequestInputDto

	if err := ctx.BodyParser(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if err := utils.ValidateStruct(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	_, err := trh.travelRequestService.Create(ctx.Context(), &input)

	if err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx,http.StatusCreated, "success created travel request", nil)
}