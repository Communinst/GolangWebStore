package authToken

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			secret := os.Getenv("AUTHORIZATION_TOKEN_SECRET")
			if secret == "" {
				http.Error(w, "error", http.StatusInternalServerError)
				return
			}
			authHeader := r.Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := IsAuthorized(authToken, secret)
				if authorized {
					userID, err := ExtractIDFromToken(authToken, secret)
					if err != nil {
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}
					ctx := context.WithValue(r.Context(), "user-id", userID)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, "Not authorized", http.StatusUnauthorized)
		})
	}
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("ID not found in token")
	}

	return id, nil
}
