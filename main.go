package main

import (
	"log"
	"net/http"

	"github.com/WalterPaes/ManUtdApi/player"
	"github.com/WalterPaes/ManUtdApi/team"
)

func main() {
	http.HandleFunc("/", team.Handler)
	http.HandleFunc("/player/", player.Handler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
