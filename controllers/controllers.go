package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CodeAkio/personalities-catalog-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
