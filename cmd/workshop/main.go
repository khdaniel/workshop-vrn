package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/handler"
	"workshop/internal/handler/config"
)

func main() {

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("Starting server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Shutting server down")
}
