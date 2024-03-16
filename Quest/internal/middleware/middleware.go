package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
		log.Printf("%s %s", req.Method, req.RequestURI)
	}
}
