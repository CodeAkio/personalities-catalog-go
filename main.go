package main

import (
	"fmt"

	"github.com/CodeAkio/personalities-catalog-go/database"
	"github.com/CodeAkio/personalities-catalog-go/routes"
)

func main() {
	database.ConnectDb()

	fmt.Println("🚀 Starting server at: http://localhost:8000")
	routes.HandleRequest()
}
