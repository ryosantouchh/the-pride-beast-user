package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/ryosantouchh/the-pride-beast-user/internal/config"
	"github.com/ryosantouchh/the-pride-beast-user/internal/infrastructure/server"
)

func main() {
	config := config.Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		panic("error loading env file")
	}

	app := server.NewServerApp()
	app.Start(config.Port)
}
