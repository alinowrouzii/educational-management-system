package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alinowrouzii/educational-management-system/controllers"
	"github.com/alinowrouzii/educational-management-system/token"
)

func TokenMiddleware(next http.Handler, jwt *token.JWTMaker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}

		payload, err := jwt.VerifyToken(authHeader[1])

		if err != nil {
			controllers.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// fmt.Println("payload inside middleware", payload)
		rcopy := r.WithContext(context.WithValue(r.Context(), "payload", payload))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, rcopy)
	})
}
