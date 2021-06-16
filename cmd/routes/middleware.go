package routes

import (
	"log"
	"net/http"
)

func withAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check request session for auth
		log.Println("Is Auth")
		f(w, r)
	}
}
