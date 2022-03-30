package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greyhands2/goMysqlBookApi/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))

}
