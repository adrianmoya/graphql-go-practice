package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/adrianmoya/graphql-go-practice/jwt"
	jwt2 "github.com/dgrijalva/jwt-go"
)

type contextKey string

const contextKeyUser = contextKey("user")

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.VerifyToken(r)
		if err == nil {
			user := token.Claims.(jwt2.MapClaims)["user_id"]
			ctx := context.WithValue(r.Context(), contextKeyUser, user)
			r = r.WithContext(ctx)
		} else {
			log.Println("Auth middleware err: ", err)
		}
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) string {
	user, _ := ctx.Value(contextKeyUser).(string)
	return user
}
