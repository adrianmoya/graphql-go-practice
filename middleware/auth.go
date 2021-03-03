package middleware

import (
	"log"
	"net/http"

	"github.com/adrianmoya/graphql-go-practice/jwt"
	jwt2 "github.com/dgrijalva/jwt-go"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.VerifyToken(r)
		if err != nil {
			log.Println("Error getting authentication token")
		}
		user := token.Claims.(jwt2.MapClaims)["user_id"]
		log.Println(user)
		next.ServeHTTP(w, r)
	})
}
