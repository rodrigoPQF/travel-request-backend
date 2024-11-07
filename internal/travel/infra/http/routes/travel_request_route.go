package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/http/handlers"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/services"
)

type travelRequestRoutes struct {
	travelRequest *services.TravelRequestService
	app *fiber.App
}

func NewTravelRequestRouteProvider(server *fiber.App, travelRequest *services.TravelRequestService) *travelRequestRoutes {
	return &travelRequestRoutes{
		travelRequest: travelRequest,
		app: server,
	}
}

func (trr *travelRequestRoutes) Routes() {
	travelRequestHandler := handlers.NewTravelRequestHandler(*trr.travelRequest)
	travelRoute := trr.app.Group("/v1/travel")
	travelRoute.Get("/request/all", travelRequestHandler.ListAllTravelRequest)
	travelRoute.Post("/request/", travelRequestHandler.CreateTravelRequest)
	travelRoute.Get("/request/:id", travelRequestHandler.ListTravelRequest)
	travelRoute.Put("/request/:id", travelRequestHandler.UpdateTravelRequest)
}