package middleware

import (
	"context"
	"net/http"

	"github.com/adrianmoya/graphql-go-practice/jwt"
	jwt2 "github.com/dgrijalva/jwt-go"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.VerifyToken(r)
		if err == nil {
			user := token.Claims.(jwt2.MapClaims)["user_id"]
			ctx := context.WithValue(r.Context(), "username", user)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *string {
	raw, _ := ctx.Value("username").(*string)
	return raw
}
