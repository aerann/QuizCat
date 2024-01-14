package main

import (
	"fmt"
	"net/http"

	"quizcat-server/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting quizcat-server")

	err := godotenv.Load()
    if err != nil {
        log.Error("Error loading .env file")
    }

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}

