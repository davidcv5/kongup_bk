package main

import (
	"log"

	"github.com/davidcv5/kongup/kongdeck"
	"github.com/davidcv5/kongup/kongfig"
)

func main() {
	kongfig, err := kongfig.GetKongfigFromFile("dev.json")
	if err != nil {
		log.Fatal(err)
	}
	err = kongdeck.KongfigToDeck(kongfig, "newKong.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
