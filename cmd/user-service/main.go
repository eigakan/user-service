package main

import (
	"fmt"

	nats_client "github.com/eigakan/nats-shared/client"
	"github.com/eigakan/user-service/config"
	"github.com/eigakan/user-service/internal/db"
	"github.com/eigakan/user-service/internal/handler"
	"github.com/eigakan/user-service/internal/model"
	"github.com/eigakan/user-service/internal/repository"
)

func main() {
	config := config.Load()

	db, err := db.Init(config.Db)

	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v\n", err))
	}

	nc, err := nats_client.NewClient(config.Nats.Host, config.Nats.Port)

	if err != nil {
		panic(err)
	}

	defer nc.Drain()

	if config.Env == "dev" {
		db.AutoMigrate(&model.User{})
	}
	userRepo := repository.NewUserRepository(db)
	userHandlers := handler.NewUserHandlers(nc, userRepo, &config.Jwt)

	userHandlers.RegisterHandlers()

	select {}
}
