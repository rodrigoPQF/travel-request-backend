package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/services"
)

type TravelRequestHandlers interface {
	CreateTravelRequest(ctx *fiber.Ctx) error
	ListTravelRequest(ctx *fiber.Ctx) error
	UpdateTravelRequest(ctx *fiber.Ctx) error
	ListAllTravelRequest(ctx *fiber.Ctx) error
}

type travelRequestHandler struct {
	travelRequestService services.TravelRequestService
}

func NewTravelRequestHandler(travelRequestService services.TravelRequestService) TravelRequestHandlers {
	return &travelRequestHandler{
		travelRequestService: travelRequestService,
	}
}