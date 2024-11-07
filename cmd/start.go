package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/rodrigoPQF/travel-request-backend/internal/config"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/database"
	"github.com/rodrigoPQF/travel-request-backend/internal/infra/http/server"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/repositories"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/services"
)

func StartApp(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[TravelRequestApp] ")


	if err := config.InitAppConfig(); err != nil {
		log.Fatalln("ocurred an error on initalizing app config", err)
	}

	conn, err := database.NewPostgresDb()
	if err != nil {
		log.Fatalln("ocurred an error on initalizing database config", err)
	}

	db := repositories.TravelRequestDB{PostgresDB: conn}

	if err := db.Migrations(); err != nil {
		log.Fatalln("ocurred an error on running migrations", err)
	}

	travelRequestService := services.NewTravelRequestService(&db)

	httpServer := server.NewHttpServer(&travelRequestService)

	sigInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(sigInterruptChannel, os.Interrupt)


	go func() {
		err = httpServer.Start(config.GetEnvs().PORT)
		if err != nil {
			log.Printf("an error ocurred on initializing http server %v", err)
		}
	}()


	<-sigInterruptChannel

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()


	<-ctx.Done()
}