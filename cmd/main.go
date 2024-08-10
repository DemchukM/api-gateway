package cmd

import (
	"github.com/DemchukM/api-gateway/internal/infrastructure"
	"github.com/DemchukM/api-gateway/internal/interfaces"
	"github.com/DemchukM/api-gateway/internal/usecase"
	"net/http"
)

func main() {
	userRepo := infrastructure.NewInMemoryUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := interfaces.NewUserHandler(userUseCase)

	http.HandleFunc("/user", userHandler.GetUser)
	http.HandleFunc("/user/create", userHandler.CreateUser)

	http.ListenAndServe(":8080", nil)
}
