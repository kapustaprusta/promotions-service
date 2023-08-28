package main

import (
	"log"

	serviceConfig "github.com/kapustaprusta/promotions-service/v2/internal/config"
	"github.com/kapustaprusta/promotions-service/v2/internal/repository"
	"github.com/kapustaprusta/promotions-service/v2/internal/services"
	httpTransport "github.com/kapustaprusta/promotions-service/v2/internal/transport/http_transport"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// TODO: add graceful shutdown
func run() error {
	config, err := serviceConfig.Read()
	if err != nil {
		return err
	}

	promotionRepository := repository.NewInMemPromotionRepository()
	promotionService := services.NewPromotionService(promotionRepository)

	log.Printf("Starting HTTP server on %s", config.BindAddr)

	if err := httpTransport.Start(config, promotionService); err != nil {
		log.Fatalf("HTTP server start error: %v", err)
	}

	log.Printf("Stopping HTTP server on %s", config.BindAddr)

	return nil
}
