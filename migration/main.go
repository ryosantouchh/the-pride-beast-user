package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/ryosantouchh/the-pride-beast-user/internal/config"
	"github.com/ryosantouchh/the-pride-beast-user/internal/core/domain/entity"
	"github.com/ryosantouchh/the-pride-beast-user/internal/infrastructure/database"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		panic("Error loading env file")
	}

	// DANGER : only use for migrate database !!
	err := database.Connect().AutoMigrate(
		&entity.Role{},
		&entity.User{},
	)
	if err != nil {
		fmt.Printf("Error Migration: %v\n", err)
	}
}
