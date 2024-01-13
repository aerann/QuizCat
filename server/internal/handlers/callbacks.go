package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"quizcat-server/api"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GenerateCards(w http.ResponseWriter, r *http.Request) {
	var params = api.GenerateCardsParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(params.Text, params.Questions)

	var response = api.GenerateCardsResponse{
		Cards: []api.Card{},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		return
	}
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}