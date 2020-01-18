package team

import (
	"encoding/json"
	"github.com/WalterPaes/ManUtdApi/pkg/domains/player"
	"net/http"
)

type Service interface {
	GetPlayer(name string) ([]byte, error, int)
	GetAllPlayers() ([]byte, error, int)
}

// Team Data ...
type Team struct {
	Player []player.Player `json:"player"`
	svc    Service         `json:"-"`
}

func New(service Service) *Team {
	return &Team{svc: service}
}

func (teamPlayer *Team) GetPlayer(name ...string) (string, int) {
	var data []byte
	var err error
	var status int

	if len(name) > 0 {
		data, err, status = teamPlayer.svc.GetPlayer(name[0])
	} else {
		data, err, status = teamPlayer.svc.GetAllPlayers()
	}

	if err != nil {
		return err.Error(), status
	}

	err = json.Unmarshal(data, teamPlayer)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	response, err := json.Marshal(teamPlayer)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	return string(response), status
}