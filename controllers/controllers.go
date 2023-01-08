package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CodeAkio/personalities-catalog-go/database"
	"github.com/CodeAkio/personalities-catalog-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}
