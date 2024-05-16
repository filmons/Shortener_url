package routers

import (
	// "net/http"
	"github.com/filmons/go-url-backend/interfaces/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter(urlController *controllers.UrlController, userController *controllers.UserController) *mux.Router {
    router := mux.NewRouter()
    // URL routes
    router.HandleFunc("/urls", urlController.GetAllUrls).Methods("GET")
    router.HandleFunc("/urls/{id}", urlController.GetUrlByID).Methods("GET")
    router.HandleFunc("/urls", urlController.CreateUrl).Methods("POST")
    router.HandleFunc("/urls/{id}", urlController.UpdateUrl).Methods("PATCH")
    router.HandleFunc("/urls/{id}", urlController.DeleteUrl).Methods("DELETE")
    // User routes
    router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
    
    return router
}





