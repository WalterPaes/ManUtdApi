package team

import (
	"github.com/WalterPaes/ManUtdApi/pkg/domains/player"
)

// Team Data ...
type Team struct {
	Player []player.Player `json:"player"`
}
