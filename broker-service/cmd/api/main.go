package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PORT string = "80"
)

type App struct {
	newClient *HTTPClient
}

func main() {
	newClient := NewHTTPClient()
	app := App{newClient: newClient}

	log.Printf("Server Running at %v \n", PORT)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
