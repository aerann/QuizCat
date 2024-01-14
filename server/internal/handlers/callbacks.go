package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"unicode"

	"quizcat-server/api"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	log "github.com/sirupsen/logrus"
)

func GenerateCards(w http.ResponseWriter, r *http.Request) {

	var params = api.GenerateCardsParams{}
	decode_err := json.NewDecoder(r.Body).Decode(&params)
	if decode_err != nil {
		log.Error(decode_err)
	}
	defer r.Body.Close()

	var cohereAPIresult string = cohereRequest(params.Text)
	// fmt.Println("COHERE RESULT: ", cohereAPIresult)

	var cards []api.Card = []api.Card{}
	parseQuestionAnswer(cohereAPIresult, &cards)

	var response = api.GenerateCardsResponse{
		Cards: cards,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		return
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func cohereRequest(input string) string {
	var auth_key string = os.Getenv("COHERE_PROD")
	client := cohereclient.NewClient(cohereclient.WithToken(auth_key))

	// Request inputs!
	length, err1 := cohere.NewSummarizeRequestLengthFromString("long")
	format, err2 := cohere.NewSummarizeRequestFormatFromString("paragraph")
	model := "command"
	extractiveness := cohere.SummarizeRequestExtractiveness("low")
	temp := 2.0
	additional_command := `Generate a summary consisting of Questions and Answers from the input that can be used as study material. Try to create as many question and answer pairs as possible, but also keep answers as concise as possible. Format as such:
	1. Question 1
	ANSWER: Answer 1
	
	2. Question 2
	ANSWER: Answer 2

	...
	`

	req := cohere.SummarizeRequest{
		Text:              input,
		Length:            &length,
		Format:            &format,
		Model:             &model,
		Extractiveness:    &extractiveness,
		Temperature:       &temp,
		AdditionalCommand: &additional_command,
	}

	response, err3 := client.Summarize(context.TODO(), &req)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Error(err1)
		log.Error(err2)
		log.Error(err3)
	}

	// fmt.Println("COHERE: ", *response.Summary)
	return *response.Summary
}

func parseQuestionAnswer(input string, ret *[]api.Card) {
	var currentCard api.Card
	var isQuestion bool

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "ANSWER:") || strings.HasPrefix(line, "Answer:") {
			currentCard.Answer = strings.TrimPrefix(line, "ANSWER: ")
			currentCard.Answer = strings.TrimPrefix(currentCard.Answer, "Answer: ")
			*ret = append(*ret, currentCard)
			currentCard = api.Card{}
			isQuestion = false
		} else if strings.HasSuffix(line, "?") || unicode.IsDigit([]rune(line)[0]) {
			if isQuestion {
				*ret = append(*ret, currentCard) // Save the previous Q&A pair if a new question starts
				currentCard = api.Card{}
			}
			currentCard.Question = line
			isQuestion = true
		} else if isQuestion {
			currentCard.Question += " " + line
		} else {
			currentCard.Answer += " " + line
		}
	}

	// Add the last Q&A pair if it exists
	if currentCard.Question != "" && currentCard.Answer != "" {
		*ret = append(*ret, currentCard)
	}
}
