package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/http/handlers"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/utils"
)

// ListAllTravelRequest godoc
// @Summary Retrieves a list of all travel requests
// @Description Fetches all travel requests based on query parameters
// @Tags Travel Request
// @Accept json
// @Produce json
// @Param status query string false "Status filter" Enums(APPROVED, CANCELED)
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} handlers.BaseResponseHandler{payload=[]dtos.TravelRequestDto}
// @Failure 400 {object} handlers.BaseResponseHandler
// @Router /request/all [get]
func (trh *travelRequestHandler) ListAllTravelRequest(ctx *fiber.Ctx) error {
	var input dtos.GetAllTravelRequestInputDto

	page := ctx.QueryInt("page", 1)    
	limit := ctx.QueryInt("limit", 10) 

	input.Limit = limit
	input.Page = page

	if err := ctx.QueryParser(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if err := utils.ValidateStruct(&input); err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := trh.travelRequestService.GetAll(ctx.Context(), &input)
	if err != nil {
		return handlers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx,http.StatusOK, "success find travel request", output)
}