package main

import (
	"library/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("server started at port :9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
