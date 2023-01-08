package main

import (
	"fmt"

	"github.com/CodeAkio/personalities-catalog-go/models"
	"github.com/CodeAkio/personalities-catalog-go/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	fmt.Println("🚀 Starting server at: http://localhost:8000")
	routes.HandleRequest()
}
