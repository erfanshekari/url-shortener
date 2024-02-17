package cli

import (
	"log"

	"github.com/erfanshekari/url-shortener/config"
)

func CliHandler(args []string) {
	if len(args) < 2 {
		log.Fatal("You must provide config file.")
	}

	switch args[1] {
	case "-c":
		if len(args) < 3 {
			log.Fatal("You must provide config file.")
		}
		conf := config.ReadConfig(args[2])
		config.SetConfig(*conf)
	default:
		printHelp()
	}
}
