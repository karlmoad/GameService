package game

import (
	"encoding/json"
	"github.com/gbrlsnchs/jwt/v2"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type GameResponse struct {
	User 	string 	`json:"user"`
	Card 	playerCard `json:"card"`
}

func GameDrawHandler(w http.ResponseWriter, r *http.Request) {
	token := context.Get(r, "TOKEN").(jwt.JWT)
	params := mux.Vars(r)

	count := 1

	if val, ok := params["count"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			count = num
		}
	}

	game, err := NewGame("POWERBALL")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Error"))
		return
	}

	card, err := game.Draw(count)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte("Internal Error"))
		return
	}


	resp := &GameResponse{User:token.Subject, Card: *card}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	return

}