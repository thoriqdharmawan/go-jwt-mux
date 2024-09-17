package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thoriqdharmawan/go-jwt-mux/config"
	"github.com/thoriqdharmawan/go-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")

		if err != nil {
			switch err {
			case http.ErrNoCookie:
				response := map[string]string{"error": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			default:
				response := map[string]string{"error": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}

		}

		tokenString := c.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			response := map[string]string{"error": err.Error()}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := map[string]string{"error": "Unauthorized, invalid token"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
