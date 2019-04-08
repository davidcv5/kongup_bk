package kongfig

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// GetKongfigFromFile loads kong configuration generated with kongfig
func GetKongfigFromFile(filename string) (*Config, error) {
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}
	fileContent, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

func readFile(kongStateFile string) (*Config, error) {

	var s Config
	b, err := ioutil.ReadFile(kongStateFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
