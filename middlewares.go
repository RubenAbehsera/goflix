package main

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"log"
	"net/http"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

// Must be logged to access
func (s *server) loggedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := jwtmiddleware.New(jwtmiddleware.Options{ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_APP_KEY), nil
		}})

		j.HandlerWithNext(w, r, next)
	}
}
