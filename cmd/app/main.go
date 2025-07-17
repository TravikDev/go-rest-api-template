package main

import (
	"fmt"
	"log"

	"go-rest-api-template/internal/config"
	"go-rest-api-template/internal/db"
	"go-rest-api-template/internal/handler"
	"go-rest-api-template/internal/repository"
	"go-rest-api-template/internal/server"
)

func main() {
	cfg := config.Load()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	database, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	userRepo := repository.NewUserRepository(database)
	charRepo := repository.NewCharacterRepository(database)

	userHandler := handler.NewUserHandler(userRepo)
	authHandler := handler.NewAuthHandler(userRepo, cfg.JWTSecret)
	charHandler := handler.NewCharacterHandler(charRepo)

	srv := server.New(userHandler, authHandler, charHandler, cfg.ServerPort, cfg.JWTSecret)

	log.Printf("starting server on %s", cfg.ServerPort)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
