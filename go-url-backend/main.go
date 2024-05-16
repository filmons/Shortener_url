package main

import (
	"log"
	"net/http"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/filmons/go-url-backend/application/services"
	"github.com/filmons/go-url-backend/infrastructure/repository"
	"github.com/filmons/go-url-backend/interfaces/controllers"
	"github.com/filmons/go-url-backend/interfaces/routers"
	"github.com/filmons/go-url-backend/pkg/config"
	"github.com/filmons/go-url-backend/pkg/utils"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	// Repository and service setup for todos
	urlRepo := repository.NewUrlRepository(db)
	urlService := services.NewUrlService(urlRepo)
	urlController := controllers.NewUrlController(urlService)

	userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

	// Router setup
	router := routers.SetupRouter(urlController, userController)

	// Start the server
	log.Println("Starting server on :8080")
	utils.Info("Server started")
	// Start the HTTP server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
