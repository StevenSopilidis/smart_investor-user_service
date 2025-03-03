package main

import (
	"log"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/db/postgres"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("---> Could not load config: ", err)
	}

	_, err = postgres.NewPostgresRepo(config)
	if err != nil {
		log.Fatal("---> Could not create postgres repo: ", err)
	}
}
