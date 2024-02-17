package main

import (
	"log"
	"os"

	"github.com/erfanshekari/url-shortener/cli"
	"github.com/erfanshekari/url-shortener/db"
	"github.com/erfanshekari/url-shortener/server"
)

func main() {

	cli.CliHandler(os.Args)

	db := db.GetInstance()

	if err := db.Start(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server := server.Server{}

	server.Init()

	server.Listen()
}
