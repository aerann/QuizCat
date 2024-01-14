package api

type Card struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type GenerateCardsParams struct {
	Text string `json:"text"`
}

type GenerateCardsResponse struct {
	Cards []Card `json:"cards"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}