package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"urlShortener/internal/cache"
	"urlShortener/internal/config"
	"urlShortener/internal/database"
	"urlShortener/internal/models"
	"urlShortener/internal/transport/rest"
)

func main() {
	config.LoadConfig()

	ErrLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	db, err := database.ConnectDB()
	if err != nil {
		ErrLogger.Fatal(err)
	}

	redis, err := cache.NewRedisClient(context.Background())
	if err != nil {
		ErrLogger.Fatal(err)
	}

	urlModel := models.NewUrlModelPostgres(db)
	Shortener := rest.NewShortener(redis, urlModel, ErrLogger, InfoLogger)

	InfoLogger.Println("Запуск сервера:", config.AppConfig.ServerAddres)

	server := http.Server{
		Handler:  Shortener.SetRoutes(),
		ErrorLog: ErrLogger,
		Addr:     config.AppConfig.ServerAddres,
	}

	if err := server.ListenAndServe(); err != nil {
		ErrLogger.Fatal(err)
	}
}
