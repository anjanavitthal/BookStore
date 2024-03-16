package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anjanavitthal/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started at 8000")

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
