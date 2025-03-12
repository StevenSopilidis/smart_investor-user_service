package gapi

import (
	"log"
	"os"
	"testing"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/db/postgres"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/utils"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/services"
)

var server *Server

func TestMain(m *testing.M) {
	config := config.Config{
		DBConnection: "postgresql://root:secret@localhost:5432/smart_investor_user_test_db?sslmode=disable",
	}

	postgresRepo, err := postgres.NewPostgresRepo(config)
	if err != nil {
		log.Fatal("Could not create postgres repo: ", err)
	}

	generator := utils.NewStringGenerator()

	user_service, err := services.NewUserService(postgresRepo, generator, 10)
	if err != nil {
		log.Fatal("Could not create user_service: ", err)
	}

	server = NewServer(user_service, config)

	code := m.Run()
	os.Exit(code)
}
