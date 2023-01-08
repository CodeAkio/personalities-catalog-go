package routes

import (
	"log"
	"net/http"

	"github.com/CodeAkio/personalities-catalog-go/controllers"
	"github.com/CodeAkio/personalities-catalog-go/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
