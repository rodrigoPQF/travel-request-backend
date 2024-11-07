package server

import (
	"context"
	"fmt"
	"log"

	_ "github.com/rodrigoPQF/travel-request-backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/http/routes"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/services"
)

type HttpServerInterface interface {
	ShutDown(ctx context.Context) error
	Start(port string) error
}

type HttpServer struct {
	app *fiber.App
	travelRequest *services.TravelRequestService
}

func NewHttpServer(travelRequest *services.TravelRequestService) HttpServerInterface {
	return &HttpServer{
		app: fiber.New(),
		travelRequest: travelRequest,
	}
}

func (hs *HttpServer) ShutDown(ctx context.Context) error {
	return hs.app.ShutdownWithContext(ctx)
}

func (hs *HttpServer) Start(port string) error {
	hs.app.Use(cors.New())
	hs.app.Use(logger.New(logger.Config{TimeZone:"America/Sao_Paulo"}))
	hs.app.Get("/swagger/*", swagger.HandlerDefault)
	routes.NewTravelRequestRouteProvider(hs.app, hs.travelRequest).Routes()

	log.Println("initializing http server")
	
	return hs.app.Listen(fmt.Sprintf(":%s", port))
}