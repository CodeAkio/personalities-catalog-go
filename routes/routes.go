package routes

import (
	"log"
	"net/http"

	"github.com/CodeAkio/personalities-catalog-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", r))
}
