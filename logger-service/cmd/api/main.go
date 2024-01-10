package main

import (
	"fmt"
	"log"
	"log-service/data"
	"net/http"
)

const (
	PORT string = "80"
)

type App struct {
	Models data.Models
}

func main() {
	app := App{}

	log.Printf("Server Running at %v \n", PORT)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
