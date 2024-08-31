package main

import (
	"log"

	"github.com/guluzadehh/kode_test/cmd/api"
)

func main() {
	apiServer := api.NewAPIServer(":8000", nil)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
