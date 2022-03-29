package routes

import (
	"github.com/gorilla/mux"
	"github.com/greyhands2/goMysqlBookApi/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/book/", controllers.getBooks).Methods("GET")

	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

func main() {

}
