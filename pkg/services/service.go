package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TeamSvc struct{}

var ErrOccurred = errors.New("an error occurred")

func (_ TeamSvc) GetPlayer(name string) ([]byte, error, int) {
	name = strings.Replace(name, " ", "%20", -1)
	team := strings.Replace("Man United", " ", "%20", -1)

	url := fmt.Sprintf(
		"https://www.thesportsdb.com/api/v1/json/1/searchplayers.php?t=%s&p=%s",
		team,
		name)

	res, err := http.Get(url)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, ErrOccurred, res.StatusCode
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return body, nil, res.StatusCode
}

func (_ TeamSvc) GetAllPlayers() ([]byte, error, int) {
	team := strings.Replace("Man United", " ", "%20", -1)

	url := fmt.Sprintf(
		"https://www.thesportsdb.com/api/v1/json/1/searchplayers.php?t=%s",
		team)

	res, err := http.Get(url)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, ErrOccurred, res.StatusCode
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return body, nil, res.StatusCode
}