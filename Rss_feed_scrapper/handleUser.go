package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abhiraj-ku/rss_feed_scrapper/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErro(w, 400, fmt.Sprint("Error handling JSON:", err))
		return
	}
	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams)
	respondWithJson(w, 200, struct{}{})
}
