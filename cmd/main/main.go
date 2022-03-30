package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greyhands2/goMysqlBookApi/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	fmt.Println("Server running at 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", router))

}
