package main

import (
	"github.com/tnyidea/react-admin-dataprovider/go/httpserver"
	"log"
)

func main() {
	server, err := httpserver.NewHttpServer(8080)
	if err != nil {
		log.Fatal("fatal: error initializing server: ", err)
	}
	defer server.Defer()

	log.Fatal(server.ListenAndServe())
}
