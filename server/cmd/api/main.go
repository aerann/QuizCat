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
	port := 8080
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting quizcat-server on port", port)

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}

	colonPort := ":" + fmt.Sprint(port)
	err = http.ListenAndServe(colonPort, r)
	if err != nil {
		log.Error(err)
	}
}
