package main

import (
	"log"

	"github.com/alvin-wilta/tiketpedia/backend/core/config"
)

func main() {
	_, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
}
