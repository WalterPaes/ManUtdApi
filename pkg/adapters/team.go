package adapters

import (
	"fmt"
	"github.com/WalterPaes/ManUtdApi/pkg/domains/team"
	"github.com/WalterPaes/ManUtdApi/pkg/services"
	"net/http"
	"strings"
)

var svc = &services.TeamSvc{}

func GetPlayer(writer http.ResponseWriter, request *http.Request) {
	var response string
	var status int

	if request.Method == http.MethodGet {
		name := strings.TrimPrefix(request.URL.Path, "/player/")
		response, status = team.New(svc).GetPlayer(name)
	} else {
		response = "Invalid Request Method!"
		status = http.StatusMethodNotAllowed
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	fmt.Fprint(writer, response)
}

func GetAllPlayers(writer http.ResponseWriter, request *http.Request) {
	var response string
	var status int

	if request.Method == http.MethodGet {
		response, status = team.New(svc).GetPlayer()
	} else {
		response = "Invalid Request Method!"
		status = http.StatusMethodNotAllowed
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	fmt.Fprint(writer, response)
}
