package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

func GenerateJWT(userId int) (string, error) {
    secretKey := "your_secret_key" // This should be stored securely
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userId,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
