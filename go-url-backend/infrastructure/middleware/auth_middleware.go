package middleware

// package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/filmons/go-url-backend/pkg/config"
	"github.com/filmons/go-url-backend/pkg/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        claims, err := ValidateToken(tokenString)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        // You can extract user ID or other claims from the JWT token and set it in the request context
        ctx := context.WithValue(r.Context(), "userID", claims["userID"])
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}


