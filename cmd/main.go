package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/routes"
)

const port = ":8080"

func main() {
	router := routes.GetV1Router()

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Server running on port", port)
	log.Fatal(server.ListenAndServe())
}