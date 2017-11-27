package main

import (
	"log"
	"net/http"

	"github.com/tomnz/glowpher/server"
)

func main() {
	log.Println("Listening...")
	handler := server.NewHandler()
	if err := http.ListenAndServe("0.0.0.0:80", handler); err != nil {
		log.Fatal(err)
	}
}
