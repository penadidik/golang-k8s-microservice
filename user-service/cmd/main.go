package main

import (
	"net/http"

	"user-service/infrastructure"
	"user-service/internal/delivery"
	"user-service/internal/repository"
	"user-service/internal/usecase"
)

func main() {
	db := infrastructure.InitMongo()
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler := delivery.NewUserHandler(userUsecase)

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	http.ListenAndServe(":8081", nil)
}
