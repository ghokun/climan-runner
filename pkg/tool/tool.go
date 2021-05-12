package tool

import (
	"encoding/json"
	"io/ioutil"
)

type Tool struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Supports    int    `json:"supports,omitempty"`
	Latest      string `json:"latest,omitempty"`
}

var (
	Tools = map[string]Tool{}
)

func GenerateTools() (err error) {
	data, err := json.Marshal(Tools)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./docs/tools.json", data, 0644)
	return err
}
