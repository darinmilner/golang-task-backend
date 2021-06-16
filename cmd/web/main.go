package main

import (
	"log"
	"net/http"
	"todos/api/cmd/routes"
	"todos/api/postgres"
)

func main() {

	db, err := postgres.NewDB("postgresql://fakepostgrespassword")
	if err != nil {
		log.Fatal(err)
	}

	h := routes.NewServer(db)
	log.Print("Server started")
	err = http.ListenAndServe(":8080", h)
	if err != nil {
		panic(err)
	}
}
