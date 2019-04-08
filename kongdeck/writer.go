package kongdeck

import (
	"io/ioutil"

	"github.com/davidcv5/kongup/kongfig"
	"gopkg.in/yaml.v2"
)

// KongfigToDeck maps kongfig to deck and saves it to a file
func KongfigToDeck(config *kongfig.Config, filename string) error {
	file, err := fromKongfig(config)
	if err != nil {
		return err
	}
	c, err := yaml.Marshal(file)
	err = ioutil.WriteFile(filename, c, 0600)
	if err != nil {
		return err
	}
	return nil
}
