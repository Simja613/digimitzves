package main

import (
	"log"
	"net/http"

	"github.com/digimitzves/core/internal/api"
)

func main() {

	log.Println("DigiMitzves Web starting")

	api.RegisterRoutes()

	log.Println("Web server listening on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		log.Fatal(err)

	}

}
