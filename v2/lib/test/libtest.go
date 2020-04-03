package main

import (
	tonlib "github.com/mercuryoio/tonlib-go/v2"
	"log"
)

func main() {
	options, err := tonlib.ParseConfigFile("../../tonlib.config.json.example")
	if err != nil {
		log.Fatal("failed parse config error. ", err)
	}

	req := tonlib.TonInitRequest{
		Type:    "init",
		Options: *options,
	}

	_, err = tonlib.NewClient(&req, tonlib.Config{}, 60, false, 0)
	if err != nil {
		log.Fatalln("Init client error", err)
	}
}
