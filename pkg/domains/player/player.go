package player

import (
	"encoding/json"
	"net/http"
)

type Service interface {
	// Search(name string) error
	GetPlayer(name string) ([]byte, error, int)
}

// Player Data ...
type Player struct {
	Data Data    `json:"player"`
	svc  Service `json:"-"`
}

type Data []struct {
	StrPlayer      string `json:"strPlayer"`
	DateBorn       string `json:"dateBorn"`
	StrNationality string `json:"strNationality"`
	StrPosition    string `json:"strPosition"`
	StrNumber      string `json:"strNumber"`
	StrWage        string `json:"strWage"`
}

func New(service Service) *Player {
	return &Player{svc: service}
}

func (p *Player) Search(name string) (string, int) {
	data, err, status := p.svc.GetPlayer(name)
	if err != nil {
		return err.Error(), status
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	response, err := json.Marshal(p)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	return string(response), status
}
