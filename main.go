package main

import (
	"github.com/WalterPaes/ManUtdApi/pkg/adapters"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", adapters.GetAllPlayers)
	http.HandleFunc("/player/", adapters.GetPlayer)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
