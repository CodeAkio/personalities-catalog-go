package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CodeAkio/personalities-catalog-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
