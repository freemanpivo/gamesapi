package main

import (
	"log"

	"github.com/freemanpivo/games-api/internal/config"
	"github.com/freemanpivo/games-api/internal/server"
)

func main() {
	// load config files;
	cfg, err := config.AppSettings()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("starting %s on port %d", cfg.App.Name, cfg.Server.Port)

	_, err = config.AppFeatures()
	if err != nil {
		log.Fatalf("failed to load feature flags: %v", err)
	}
	log.Printf("FUNC Limiting user games by %d", config.Limit("maxGamesPerUser", 0))
	// load databases;

	// create server;
	server.Start()
	// create base route;

}
