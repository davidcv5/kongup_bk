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

// func mapUpstream(api *API) (*upstream, error) {
// 	upstream := &upstream{}
// 	u, err := url.Parse(api.Attributes.UpstreamURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid url: %s", api.Attributes.UpstreamURL)
// 	}
// 	upstream.Name = &u.Host
// 	upstream.Targets
// 	return upstream, nil
// }
