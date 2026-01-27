package main

import (
	"enriqueFcoG/cypher/internal/config"
	"enriqueFcoG/cypher/internal/di"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello worl to this new amazing app")
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Println("error reading system envs")
		}
	}
	cfg := config.Load()
	fmt.Println(cfg)
	server := di.BuildServer(cfg)
	server.Run(":" + cfg.Port)
}
