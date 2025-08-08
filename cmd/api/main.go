package main

import (
	"fmt"

	"github.com/eigakan/user-service/config"
	"github.com/eigakan/user-service/internal/db"
	"github.com/eigakan/user-service/internal/model"
	"github.com/eigakan/user-service/internal/nats"
)

func main() {
	config := config.Load()

	db, err := db.Init(config.Db)

	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
	}

	nc, err := nats.NewClient(config.Nats)
	if err != nil {
		fmt.Printf("Error connecting to NATS: %v\n", err)
	}
	defer nc.Close()

	db.AutoMigrate(&model.User{})

	select {}
}
