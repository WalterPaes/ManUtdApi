package team

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WalterPaes/ManUtdApi/player"
)

type Team struct {
	Player []player.Player
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == "GET":
		response := GetPlayers()

		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprint(writer, response)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Method not valid!")
	}
}

func GetPlayers() string {
	res, err := http.Get("https://www.thesportsdb.com/api/v1/json/1/searchplayers.php?t=Man%20United")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var players Team
	json.Unmarshal(body, &players)

	json, err := json.Marshal(players)
	if err != nil {
		panic(err)
	}

	return string(json)
}
