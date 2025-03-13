package main

import (
	"log"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/db/postgres"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/gapi"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/utils"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/services"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("---> Could not load config: ", err)
	}

	repo, err := postgres.NewPostgresRepo(config)
	if err != nil {
		log.Fatal("---> Could not create postgres repo: ", err)
	}

	stringGenerator := utils.NewStringGenerator()

	user_service, err := services.NewUserService(repo, stringGenerator, config.EmailVerificationCodeLength)
	if err != nil {
		log.Fatal("---> Could not create user_service: ", err)
	}

	server := gapi.NewServer(user_service, config)
	server.Run()
}
