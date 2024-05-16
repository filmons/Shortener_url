package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
    "github.com/filmons/go-url-backend/application/dto" // Import the dto package

	"github.com/filmons/go-url-backend/application/services"
	"github.com/filmons/go-url-backend/pkg/utils"
	"github.com/gorilla/mux" // Import gorilla/mux package for routing


)

type UrlController struct {
	Service *services.UrlService
}

func NewUrlController(s *services.UrlService) *UrlController {
	return &UrlController{
		Service: s,
	}
}

func (c *UrlController) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := c.Service.GetAllUrls()
	if err != nil {
		utils.Logger.Println("Error getting urls:", err)
		http.Error(w, "Failed to get urls", http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(urls)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *UrlController) GetUrlByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        utils.Logger.Println("Invalid URL ID:", err)
        http.Error(w, "Invalid URL ID", http.StatusBadRequest)
        return
    }

    // Corrected the method name to use Service instead of service
    url, err := c.Service.GetUrlByID(uint(id))
    if err != nil {
        utils.Logger.Println("Error getting URL by ID:", err)
        http.Error(w, "Failed to get URL by ID", http.StatusInternalServerError)
        return
    }

    response, _ := json.Marshal(url)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *UrlController) CreateUrl(w http.ResponseWriter, r *http.Request) {
    // Parse the request body into a URL DTO
    var urlDto dto.UrlDTO
    err := json.NewDecoder(r.Body).Decode(&urlDto)
    if err != nil {
        utils.Logger.Println("Error decoding request body:", err)
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Call the service to create the URL
    err = c.Service.CreateUrl(&urlDto)
    if err != nil {
        utils.Logger.Println("Error creating URL:", err)
        http.Error(w, "Failed to create URL", http.StatusInternalServerError)
        return
    }

    // Send a success response
    w.WriteHeader(http.StatusCreated)
}

func (c *UrlController) UpdateUrl(w http.ResponseWriter, r *http.Request) {
    // Parse request body to get the updated URL data
    var updatedUrl dto.UrlDTO
    err := json.NewDecoder(r.Body).Decode(&updatedUrl)
    if err != nil {
        utils.Logger.Println("Error decoding request body:", err)
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Get the URL ID from the request URL parameters
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        utils.Logger.Println("Error parsing URL parameter:", err)
        http.Error(w, "Invalid URL parameter", http.StatusBadRequest)
        return
    }

    // Call the service's Update method using the DTO
    err = c.Service.Update(uint(id), &updatedUrl)
    if err != nil {
        utils.Logger.Println("Error updating URL:", err)
        http.Error(w, "Failed to update URL", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("URL updated successfully"))
}


// func (c *UrlController) UpdateUrl(w http.ResponseWriter, r *http.Request) {
//     // Extract the URL ID from the request path parameters
//     vars := mux.Vars(r)
//     id, err := strconv.ParseUint(vars["id"], 10, 64)
//     if err != nil {
//         utils.Logger.Println("Invalid URL ID:", err)
//         http.Error(w, "Invalid URL ID", http.StatusBadRequest)
//         return
//     }

//     // Parse the request body into a URL DTO
//     var urlDto dto.UrlDTO
//     err = json.NewDecoder(r.Body).Decode(&urlDto)
//     if err != nil {
//         utils.Logger.Println("Error decoding request body:", err)
//         http.Error(w, "Failed to decode request body", http.StatusBadRequest)
//         return
//     }

//     // Call the service to update the URL
//     err = c.Service.Update(uint(id), &urlDto)
//     if err != nil {
//         utils.Logger.Println("Error updating URL:", err)
//         http.Error(w, "Failed to update URL", http.StatusInternalServerError)
//         return
//     }

//     // Send a success response
//     w.WriteHeader(http.StatusOK)
// }


func (c *UrlController) DeleteUrl(w http.ResponseWriter, r *http.Request) {
    // Extract the URL ID from the request path parameters
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        utils.Logger.Println("Invalid URL ID:", err)
        http.Error(w, "Invalid URL ID", http.StatusBadRequest)
        return
    }

    // Call the service to delete the URL
    err = c.Service.DeleteUrl(uint(id))
    if err != nil {
        utils.Logger.Println("Error deleting URL:", err)
        http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
        return
    }

    // Send a success response
    w.WriteHeader(http.StatusOK)
}
