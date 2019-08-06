package team

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/WalterPaes/ManUtdApi/player"
)

type Team struct {
	Player []player.Player `json:"player"`
}

func Handler(writer http.ResponseWriter, request *http.Request) {

	switch {
	case request.Method == "GET":
		var response string
		var err error
		player := strings.TrimPrefix(request.URL.Path, "/player/")
		if len(player) > 0 {
			response, err = GetPlayer(player)
		} else {
			response, err = GetPlayers()
		}

		if err != nil {
			panic(err.Error())
		}

		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprint(writer, response)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Method not valid!")
	}
}

func GetPlayers() (string, error) {
	res, err := http.Get("https://www.thesportsdb.com/api/v1/json/1/searchplayers.php?t=Man%20United")
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var players Team
	err = json.Unmarshal(body, &players)
	if err != nil {
		return "", err
	}

	json, err := json.Marshal(players)
	if err != nil {
		return "", err
	}

	return string(json), err
}

func GetPlayer(name string) (string, error) {

	name = strings.ReplaceAll(name, " ", "%20")
	team := strings.ReplaceAll("Man United", " ", "%20")

	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/searchplayers.php?t=%s&p=%s", team, name)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var player Team
	err = json.Unmarshal(body, &player)
	if err != nil {
		return "", err
	}

	json, err := json.Marshal(player)
	if err != nil {
		return "", err
	}

	return string(json), err
}
