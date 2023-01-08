package routes

import (
	"log"
	"net/http"

	"github.com/CodeAkio/personalities-catalog-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// personalities_catalog
