package main

import (
	"fmt"

	"urlShortener/internal/config"
	"urlShortener/internal/repository"
	"urlShortener/internal/service"
	"urlShortener/internal/transport/rest"
)

func main() {
	config.LoadConfig()

	repositoryShortener := repository.NewRepositoryShortenerPostgres()
	serviceShortener := service.NewServiceShortener(repositoryShortener)
	httpHandlersShortener := rest.NewHttpHandlersShortener(serviceShortener)

	httpServer := rest.NewHttpServerShortener(httpHandlersShortener)
	if err := httpServer.Start(); err != nil {
		fmt.Println(err)
	}
}
